package service

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"google.golang.org/grpc/metadata"
)

type userService struct {
	config  *config.Config
	store   UserStore
	criptor core.Encryptor
	auth    core.Auth
}

// LoginAccept implements UserService.
func (s *userService) ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error) {
	var (
		err  error
		user *core.User
	)
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("err metadata")
	}
	acceptTokens := data.Get(core.CtxAcceptToken)
	if len(acceptTokens) == 0 {
		return nil, errors.New("err token metadata")
	}
	token, err := hex.DecodeString(acceptTokens[0])
	if err != nil {
		return nil, err
	}
	accept, err := s.criptor.Decrypt([]byte(form.Code), token)
	if err != nil {
		return nil, errors.New("err code")
	}
	lifetimeStr := string(accept)
	lifetime, err := time.Parse(time.RFC3339, lifetimeStr)
	if err != nil {
		return nil, errors.New("err code")
	}
	if lifetime.UnixMilli() < time.Now().UnixMilli() {
		return nil, errors.New("err code code expires")
	}

	user, err = s.store.GetUserByEmail(ctx, form.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements UserService.
func (s *userService) GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error) {
	var (
		err  error
		user *core.User
	)
	user, err = s.store.GetUserByEmail(ctx, form.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		_, err = s.store.AddUser(ctx, &core.User{})
		if err != nil {
			return nil, err
		}
	}
	// TODO GENERATED CODE
	const lifetime = 1
	codeLifetime := time.Now().Add(time.Minute * lifetime).Format(time.RFC3339)
	var codeToEmail = "1234"
	token, err := s.criptor.Encrypt([]byte(codeToEmail), []byte(codeLifetime))
	if err != nil {
		return nil, err
	}
	// TODO SEND TO EMAIL

	// test, _ := s.criptor.Decrypt([]byte(codeToEmail), token)
	// log.Println(test)
	return token, nil
}

func NewUserService(conf config.Config, store UserStore, cr core.Encryptor) UserService {
	return &userService{
		config:  &conf,
		store:   store,
		criptor: cr,
		auth:    core.NewAuth(conf),
	}
}
