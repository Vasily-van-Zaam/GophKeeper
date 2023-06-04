package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type userService struct {
}

// GetSecretToken implements UserService.
func (*userService) GetSecretToken(ctx context.Context, userID string) ([]byte, error) {
	panic("unimplemented")
}

// Login implements UserService.
func (*userService) Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	panic("unimplemented")
}

// Registration implements UserService.
func (*userService) Registration(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	panic("unimplemented")
}

func NewUserService() UserService {
	return &userService{}
}
