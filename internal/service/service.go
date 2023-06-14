// The Service package is main business logic
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package service

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/golang-jwt/jwt"

	"google.golang.org/grpc/metadata"
)

type Store interface {
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
}

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	GetSecretToken(ctx context.Context, userID string) ([]byte, error)
}

type Service interface {
	GetData(ctx context.Context, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	SearchData(ctx context.Context, search string, types ...string) ([]*core.ManagerData, error)
}

type UserService interface {
	GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error)
	ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error)
}

func (s *service) handlerAuth(ctx context.Context, user *core.User) (*core.AuthToken, error) {
	data, ok := metadata.FromIncomingContext(ctx)
	// md := ctx.Value("client_version")
	// log.Println(md)
	if !ok {
		return nil, errors.New("err metadata")
	}
	vesion := data.Get("client_version")
	if len(vesion) == 0 {
		return nil, errors.New("err metadata client version")
	}
	claims := jwt.MapClaims{
		"email":  user.Email,
		"userID": fmt.Sprint(user.ID),
		"exp":    time.Now().Add(time.Hour * time.Duration(s.config.Server().Expires())).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(s.config.Server().SecretKey(vesion[0])))
	if err != nil {
		return nil, err
	}
	return &core.AuthToken{
		Access: []byte(ts),
	}, nil
}
