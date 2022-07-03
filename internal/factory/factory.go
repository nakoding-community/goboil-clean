package factory

import (
	"github.com/nakoding-community/goboil-clean/internal/factory/repository"
	"github.com/nakoding-community/goboil-clean/internal/factory/usecase"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
}

func Init() Factory {
	f := Factory{}
	f.Repository = repository.Init()
	f.Usecase = usecase.Init(f.Repository)

	return f
}
