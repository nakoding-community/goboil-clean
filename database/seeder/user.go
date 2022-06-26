package seeder

import (
	"log"

	"github.com/google/uuid"
	"github.com/nakoding-community/goboil-clean/internal/abstraction"
	"github.com/nakoding-community/goboil-clean/internal/model"

	"gorm.io/gorm"
)

func UserTableSeeder(conn *gorm.DB) {
	trx := conn.Begin()

	if err := trx.Create(&model.UserEntityModel{
		Entity: abstraction.Entity{
			ID: uuid.New(),
		},
		UserEntity: model.UserEntity{
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: "admin",
		},
		Context: Context,
	}).Error; err != nil {
		trx.Rollback()
		log.Println(err.Error())
		return
	}

	if err := trx.Commit().Error; err != nil {
		log.Println(err.Error())
	}
}
