package factory

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

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection(constant.DB_GOBOIL_CLEAN)
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
}
