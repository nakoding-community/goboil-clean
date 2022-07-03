package user

import (
	"fmt"

	"github.com/nakoding-community/goboil-clean/internal/abstraction"
	"github.com/nakoding-community/goboil-clean/internal/dto"
	"github.com/nakoding-community/goboil-clean/internal/factory"
	"github.com/nakoding-community/goboil-clean/internal/middleware"
	"github.com/nakoding-community/goboil-clean/internal/model"
	"github.com/nakoding-community/goboil-clean/pkg/constant"
	res "github.com/nakoding-community/goboil-clean/pkg/util/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	Factory factory.Factory
}

func NewHandler(f factory.Factory) *handler {
	return &handler{f}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get user
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @param request query abstraction.SearchGetRequest true "request query"
// @Param name query string false "name"
// @Success 200 {object} dto.SearchGetResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /users [get]
func (h *handler) Get(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(abstraction.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	abstraction.BindFilterSort[model.UserEntity](c, model.UserEntity{}, "users", payload)

	result, err := h.Factory.Usecase.User.Find(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get users success", &result.PaginationInfo).Send(c)
}

// Get user by id
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /users/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.User.FindByID(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create user
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /users [post]
func (h *handler) Create(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Create(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update user
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param request body dto.UpdateUserRequest true "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /users/{id} [put]
func (h *handler) Update(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.UpdateUserRequest)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, fmt.Errorf("id must be uuid")).Send(c)
	}
	payload.ID = id
	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Update(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete user
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /users/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	cc := c.Request().Context().Value(constant.CONTEXT_KEY).(*abstraction.Context)

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Delete(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
