package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type userService struct {
}

// RegistrationAccept implements UserService.
func (*userService) RegistrationAccept(ctx context.Context, form *core.LoginForm) error {
	panic("unimplemented")
}

// Login implements UserService.
func (*userService) Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	panic("unimplemented")
}

// Registration implements UserService.
func (*userService) Registration(ctx context.Context, form *core.LoginForm) (*string, error) {
	panic("unimplemented")
}

func NewUserService() UserService {
	return &userService{}
}
