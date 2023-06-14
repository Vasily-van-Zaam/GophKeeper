package repository

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type localStore interface {
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
}

type local struct {
	store  localStore
	config config.Config
	auth   core.Auth
}

// AddData implements localStore.
func (*local) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// ChangeData implements localStore.
func (*local) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// SearchData implements localStore.
func (*local) SearchData(ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// Close implements Local.
func (l *local) Close() error {
	return l.store.Close()
}

// GetData implements Local.
func (*local) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	return nil, nil
}

func NewLocal(conf config.Config, store localStore) localStore {
	return &local{
		store:  store,
		config: conf,
		auth:   core.NewAuth(conf),
	}
}
