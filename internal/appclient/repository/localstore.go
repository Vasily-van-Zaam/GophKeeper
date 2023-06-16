package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

// localstore Implements.
type localStore interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)

	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context) (*core.ManagerData, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
}

type Local interface {
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context) (*core.User, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	AddAccessData(ctx context.Context, masterPsw string, user *core.User) error
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
}

type local struct {
	store  localStore
	config config.Config
	auth   core.Auth
}

// AddAccessData implements Local.
func (l *local) AddAccessData(ctx context.Context, masterPsw string, user *core.User) error {
	manager := core.NewManager()
	manager.AddEncription(l.config.Encryptor()).
		Set().MetaData("access")
	err := manager.Set().AccessData(masterPsw, user)
	if err != nil {
		return err
	}
	mData, err := manager.ToData()
	if err != nil {
		return err
	}
	_, err = l.store.AddData(ctx, mData)
	return err
}

// GetAccessData implements localStore.
func (l *local) GetAccessData(ctx context.Context) (*core.User, error) {
	data, err := l.store.GetAccessData(ctx)
	if err != nil {
		return nil, err
	}
	// return nil, errors.New("test")
	manager := data.ToManager().AddEncription(l.config.Encryptor())

	version := l.config.Client().Version()
	log.Println("===", l.config.Server().SecretKey(version))
	d, err := manager.Get().Data(l.config.Server().SecretKey(version))
	if err != nil {
		return nil, err
	}
	var user core.User
	err = json.Unmarshal(d, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
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
func (*local) SearchData(
	ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// Close implements Local.
func (l *local) Close() error {
	return l.store.Close()
}

// GetData implements Local.
func (l *local) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	log.Println("?????", l)
	return nil, nil
}

func NewLocal(conf config.Config, store localStore) Local {
	return &local{
		store:  store,
		config: conf,
		auth:   core.NewAuth(conf),
	}
}
