package repository

import (
	"github.com/nakoding-community/goboil-clean/database"
	"github.com/nakoding-community/goboil-clean/internal/repository"
	"github.com/nakoding-community/goboil-clean/pkg/constant"
	"gorm.io/gorm"
)

type Factory struct {
	Db             *gorm.DB
	UserRepository repository.User
}

func Init() Factory {
	f := Factory{}
	f.InitDb()
	f.InitUserRepository()

	return f
}

func (f *Factory) InitDb() {
	db, err := database.GetConnection(constant.DB_GOBOIL_CLEAN)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) InitUserRepository() {
	if f.Db == nil {
		panic("Failed init repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
}
