package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	usecase "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/usecase"
	validator "gopkg.in/go-playground/validator.v9"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UUsecase usecase.IUserUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, uc usecase.IUserUseCase) {
	handler := &UserHandler{
		UUsecase: uc,
	}
	// e.GET("/articles", handler.FetchArticle)
	e.POST("/user", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

func (u *UserHandler) Store(c echo.Context) (err error) {
	var user db.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = u.UUsecase.Store(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func isRequestValid(m *db.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
