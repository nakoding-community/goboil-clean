package user

import (
	"github.com/nakoding-community/goboil-clean/internal/abstraction"
	"github.com/nakoding-community/goboil-clean/internal/dto"
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/model"
	"github.com/nakoding-community/goboil-clean/internal/repository"
	res "github.com/nakoding-community/goboil-clean/pkg/util/response"
	"github.com/nakoding-community/goboil-clean/pkg/util/str"
	"github.com/nakoding-community/goboil-clean/pkg/util/trxmanager"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.SearchGetRequest, filters []dto.Filter) (*dto.SearchGetResponse[dto.UserResponse], error)
	FindByID(ctx *abstraction.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(ctx *abstraction.Context, payload *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error)
}

type service struct {
	UserRepository repository.User
	Db             *gorm.DB
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository, f.Db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.SearchGetRequest, filters []dto.Filter) (*dto.SearchGetResponse[dto.UserResponse], error) {
	users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination, filters)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var datas []dto.UserResponse
	for _, user := range users {
		datas = append(datas, dto.UserResponse{
			UserEntityModel: user,
		})
	}

	result := new(dto.SearchGetResponse[dto.UserResponse])
	result.Datas = datas
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.UserResponse{
		UserEntityModel: data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CreateUserRequest) (*dto.UserResponse, error) {
	var email string
	if payload.Email != nil {
		email = *payload.Email
	} else {
		email = str.GenerateRandString(10) + "@gmail.com"
	}

	var (
		result *dto.UserResponse
		userID = uuid.New()
		user   = model.UserEntityModel{
			Entity: abstraction.Entity{
				ID: userID,
			},
			UserEntity: model.UserEntity{
				Name:     payload.Name,
				Email:    email,
				Password: payload.Password,
			},
			Context: ctx,
		}
	)
	var err error

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		user, err = s.UserRepository.Create(ctx, user)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.UserResponse{
		UserEntityModel: user,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	var (
		result *dto.UserResponse
		user   = model.UserEntityModel{
			UserEntity: model.UserEntity{
				Name:     payload.Name,
				Email:    payload.Email,
				Password: payload.Password,
			},
			Context: ctx,
		}
		err error
	)

	if payload.Password != "" {
		user.HashPassword()
		user.Password = ""
	}

	if payload.Token != "" {
		user.Token = payload.Token
	}

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		user, err = s.UserRepository.UpdateByID(ctx, payload.ID, user)
		if err != nil {
			return err
		}

		user, err = s.UserRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.UserResponse{
		UserEntityModel: user,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.ByIDRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse
	var data model.UserEntityModel
	var err error

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.UserRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		err = s.UserRepository.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.UserResponse{
		UserEntityModel: data,
	}

	return result, nil
}
