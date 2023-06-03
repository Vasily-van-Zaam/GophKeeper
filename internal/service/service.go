// The Service package is main business logic
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data *core.ManagerData) error
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
}
type Service interface {
	Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error)
	Registration(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error)
	GetData(ctx context.Context, types ...core.DataType) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data *core.ManagerData) error
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
}

type service struct {
	store     Store
	encriptor core.Encryptor
}

// AddData implements Service.
func (*service) AddData(ctx context.Context, data *core.ManagerData) error {
	panic("unimplemented")
}

// ChangeData implements Service.
func (*service) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// GetData implements Service.
func (*service) GetData(ctx context.Context, types ...core.DataType) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// Login implements Service.
func (*service) Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	panic("unimplemented")
}

// Registration implements Service.
func (*service) Registration(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	panic("unimplemented")
}

func New(store Store, encript core.Encryptor) Service {
	return &service{
		store:     store,
		encriptor: encript,
	}
}
