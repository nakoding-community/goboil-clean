package migration

import (
	"fmt"

	"github.com/nakoding-community/goboil-clean/database"
	"github.com/nakoding-community/goboil-clean/internal/model"
	"github.com/nakoding-community/goboil-clean/pkg/constant"
	"github.com/nakoding-community/goboil-clean/pkg/util/env"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !env.NewEnv().GetBool(constant.IS_RUN_MIGRATION) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_GOBOIL_CLEAN: &migration{
			DbModels: &[]interface{}{
				&model.UserEntityModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := database.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
