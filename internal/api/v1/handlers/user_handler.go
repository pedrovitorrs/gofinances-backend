package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	request "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/dto/request"
	response "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/dto/response"
	helper "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/helpers"
	_ "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	usecase "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/usecase"
	_ "github.com/swaggo/swag"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

type UserHandler struct {
	UUsecase usecase.IUserUseCase
}

// NewArticleHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, uc usecase.IUserUseCase) {
	handler := &UserHandler{
		UUsecase: uc,
	}
	group := e.Group("/api/v1")
	// e.GET("/articles", handler.FetchArticle)
	group.POST("/users", handler.Store)
	group.GET("/users/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags User
// Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "User Payload"
// @Success      200  {object}  sqlc.User
// @Failure 	 400  {object}  response.ResponseError
// @Router       /users [post]
func (u *UserHandler) Store(c echo.Context) (err error) {
	var user request.CreateUserRequest
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = helper.IsValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userCreated, err := u.UUsecase.Store(ctx, user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func (u *UserHandler) GetByID(c echo.Context) (err error) {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	id := int32(idParam)
	ctx := c.Request().Context()

	user, err := u.UUsecase.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
