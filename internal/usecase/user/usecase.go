package user

import (
	"context"
	"errors"

	"github.com/mohar9h/go-graphql-ws-api/internal/domain/user"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailTaken         = errors.New("email already taken")
	ErrUsernameTaken      = errors.New("username already taken")
)

func NewUseCase(userRepo user.Repository) user.UseCase {
	return &useCase{
		userRepo: userRepo,
	}
}

type useCase struct {
	userRepo user.Repository
}

func (u *useCase) Login(ctx context.Context, username string, password string) (*user.User, error) {
	panic("unimplemented")
}

func (u *useCase) Logout(ctx context.Context) error {
	panic("unimplemented")
}

func (u *useCase) RefreshToken(ctx context.Context, refreshToken string) (*user.User, error) {
	panic("unimplemented")
}

func (u *useCase) Register(ctx context.Context, username *string, mobile *string, email *string, password *string, firstName *string, lastName *string) (*user.User, error) {
	panic("unimplemented")
}

func (u *useCase) ResetPassword(ctx context.Context, email string) error {
	panic("unimplemented")
}
