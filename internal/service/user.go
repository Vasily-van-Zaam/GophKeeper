package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type userService struct {
	config  *config.Config
	store   UserStore
	criptor core.Encryptor
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

func NewUserService(conf config.Config, store UserStore, cr core.Encryptor) UserService {
	return &userService{
		config:  &conf,
		store:   store,
		criptor: cr,
	}
}
