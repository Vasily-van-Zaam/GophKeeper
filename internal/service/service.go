// The Service package is main business logic
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
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
	Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error)
	Registration(ctx context.Context, form *core.LoginForm) (*string, error)
	RegistrationAccept(ctx context.Context, form *core.LoginForm) error
}
