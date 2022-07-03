package usecase

import (
	"github.com/nakoding-community/goboil-clean/internal/factory/repository"
	"github.com/nakoding-community/goboil-clean/internal/usecase/auth"
	"github.com/nakoding-community/goboil-clean/internal/usecase/user"
)

type Factory struct {
	Auth auth.Usecase
	User user.Usecase
}

func Init(r repository.Factory) Factory {
	f := Factory{}
	f.Auth = auth.NewUsecase(r)
	f.User = user.NewUsecase(r)

	return f
}
