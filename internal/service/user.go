package service

import (
	"context"
	"errors"
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"google.golang.org/grpc/metadata"
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
func (s *userService) Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	var (
		err  error
		user *core.User
	)
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("err metadata")
	}

	data.Get("")
	user, err = s.store.GetUserByEmail(ctx, form.Email)
	if err != nil {
		return nil, err
	}
	log.Println(user, err)
	return &core.AuthToken{
		Access:  []byte("access"),
		Refresh: []byte("refresh"),
		UserKey: []byte("user_key"),
	}, nil
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
