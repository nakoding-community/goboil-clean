package repository

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nakoding-community/goboil-clean/internal/dto"
	"github.com/nakoding-community/goboil-clean/internal/model"
	res "github.com/nakoding-community/goboil-clean/pkg/util/response"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type validType interface {
	model.UserEntityModel | model.UserEntity
}

func FindAll[T validType](conn *gorm.DB, data T) (T, error) {
	err := conn.Find(&data).Error
	return data, maskError(err)
}

func FindByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) (T, error) {
	err := conn.Where("id = ?", ID).First(&data).Error
	return data, maskError(err)
}

func Create[T validType](conn *gorm.DB, data T) (T, error) {
	err := conn.Create(&data).Error
	return data, maskError(err)
}

func UpdateByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) (T, error) {
	err := conn.Model(&data).Where("id = ?", ID).Updates(&data).Error
	return data, maskError(err)
}

func DeleteByID[T validType](conn *gorm.DB, ID uuid.UUID, data T) error {
	return maskError(conn.Where("id = ?", ID).Delete(&data).Error)
}

func Deletes[T validType](conn *gorm.DB, IDs []uuid.UUID, data T) error {
	return maskError(conn.Delete(&data, IDs).Error)
}

// BuildFilterQuery
// for now questomFilter max length is 2
// index 0 for query
// index 1 for type
// example use qustomFilter:
//
//	BuildFilterQuery(query, filters, "concat(types.value, ' ',units.unit) = ?", "types.value")
//	on this example we wont to change query on field 'types.value'  from 'types.value = ?' to 'concat(types.value, ' ',units.unit) = ?
func BuildFilterQuery(query *gorm.DB, filters []dto.Filter, qustomFilters ...string) {
	var qustom = false
	if len(qustomFilters) == 2 {
		qustom = true
	}

	for _, filter := range filters {
		if qustom {
			if filter.Field == qustomFilters[1] {
				query = query.Where(qustomFilters[0], filter.Value)
				continue
			}
		}
		query.Where(filter.Field+" = ?", filter.Value)
	}
}

func GetColumnsSort[T validType](AscField, DescField []string, query *gorm.DB, data T, name string, excludes ...string) {
	var AscValids, DescValids []string

	fields := reflect.ValueOf(&data).Elem()

LoopField:
	for i := 0; i < fields.NumField(); i++ {
		column := fields.Type().Field(i).Tag.Get("json")

		for _, exclude := range excludes {
			if strings.ToLower(column) == strings.ToLower(exclude) {
				continue LoopField
			}
		}

		for _, val := range AscField {
			if val == column {
				AscValids = append(AscValids, name+"."+column)
			}
		}

		for _, val := range DescField {
			fmt.Println(val)
			if val == column {
				DescValids = append(DescValids, name+"."+column)
			}
		}
	}

	if len(AscValids) > 0 {
		columns := strings.Join(AscValids, ",")
		query = query.Order(columns + " asc")
	}

	if len(DescValids) > 0 {
		columns := strings.Join(DescValids, ",")
		query = query.Order(columns + " desc")

	}

	return
}

func maskError(err error) error {
	if err != nil {
		// not found
		if err == gorm.ErrRecordNotFound {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		if pqErr, ok := err.(*pgconn.PgError); ok {
			// duplicate data
			if pqErr.Code == "23505" {
				return res.ErrorBuilder(&res.ErrorConstant.DuplicateEntity, err)
			}
		}

		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
	}

	return nil
}
