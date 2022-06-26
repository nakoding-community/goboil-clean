package auth

import (
	"github.com/nakoding-community/goboil-clean/internal/abstraction"
	"github.com/nakoding-community/goboil-clean/internal/dto"
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/model"
	"github.com/nakoding-community/goboil-clean/internal/repository"
	res "github.com/nakoding-community/goboil-clean/pkg/util/response"
	"github.com/nakoding-community/goboil-clean/pkg/util/trxmanager"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Db         *gorm.DB
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.Db, f.UserRepository}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var data model.UserEntityModel

	data.UserEntity = payload.UserEntity

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AuthRegisterResponse{
		UserEntityModel: data,
	}

	return result, nil
}
