package pstgql

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	GetSecretToken(ctx context.Context, userID string) ([]byte, error)

	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
}

type store struct {
	config config.Config
}

// AddData implements Store.
func (*store) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// AddUser implements Store.
func (*store) AddUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// ChangeData implements Store.
func (*store) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// ChangeUser implements Store.
func (*store) ChangeUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// GetData implements Store.
func (*store) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// GetSecretToken implements Store.
func (*store) GetSecretToken(ctx context.Context, userID string) ([]byte, error) {
	panic("unimplemented")
}

// SearchData implements Store.
func (*store) SearchData(ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

func New(conf config.Config) (Store, error) {
	return &store{
		config: conf,
	}, nil
}
