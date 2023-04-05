package usecase

import (
	"bytes"
	"context"
	"crypto/sha512"
	"time"

	request "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/dto/request"
	db "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	Store(contxt context.Context, user request.CreateUserRequest) (_ db.User, err error)
}

type userUseCase struct {
	store          *db.SQLStore
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUserUseCase(store *db.SQLStore, timeout time.Duration) *userUseCase {
	return &userUseCase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (uc *userUseCase) Store(contxt context.Context, user request.CreateUserRequest) (userCreated db.User, err error) {
	ctx, cancel := context.WithTimeout(contxt, uc.contextTimeout)
	defer cancel()

	hashedInput := sha512.Sum512_256([]byte(user.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
	var passwordHashed = string(passwordHashInBytes)

	arg := db.CreateUserParams{
		Username: user.Username,
		Password: passwordHashed,
		Email:    user.Email,
	}

	return uc.store.CreateUser(ctx, arg)
}
