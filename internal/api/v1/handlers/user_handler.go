package handlers

import (
	"log"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	request "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/dto/request"
	response "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/dto/response"
	usecase "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/usecase"
)

type UserHandler struct {
	UUsecase usecase.IUserUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, uc usecase.IUserUseCase) {
	handler := &UserHandler{
		UUsecase: uc,
	}
	// e.GET("/articles", handler.FetchArticle)
	e.POST("/users", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

func (u *UserHandler) Store(c echo.Context) (err error) {
	var user request.CreateUserRequest
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userCreated, err := u.UUsecase.Store(ctx, user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func isRequestValid(req interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	log.Println(req)
	if err != nil {
		return false, err
	}
	return true, nil
}
