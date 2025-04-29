package user

import "context"

type UseCase interface {
	Register(ctx context.Context, username, mobile, email, password, firstName, lastName *string) (*User, error)
	Login(ctx context.Context, username, password string) (*User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*User, error)
	ResetPassword(ctx context.Context, email string) error
	Logout(ctx context.Context) error
}
