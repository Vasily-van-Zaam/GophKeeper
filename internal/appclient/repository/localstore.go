package repository

import (
	"context"
	"encoding/json"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

// localstore Implements.
type localStore interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	ResetUserData(ctx context.Context) error
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context) (*core.ManagerData, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
	AddTryPasword(ctx context.Context, data *core.ManagerData) error
	GetTryPasword(ctx context.Context) (*core.ManagerData, error)
	AddAccessData(ctx context.Context, data *core.ManagerData) error
}

type Local interface {
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context, masterPsw string) (*core.User, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	AddAccessData(ctx context.Context, masterPsw string, user *core.User) error
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
	ResetUserData(ctx context.Context) error
	AddTryPasword(ctx context.Context, count int) error
	GetTryPasword(ctx context.Context) (int, error)
}

type local struct {
	store  localStore
	config config.Config
	auth   core.Auth
}

// AddTryPasword implements Local.
func (l *local) AddTryPasword(ctx context.Context, count int) error {
	key := l.config.Server().SecretKey(l.config.Client().Version())
	count++
	manager := core.NewManager().AddEncription(l.config.Encryptor())
	err := manager.
		Set().MetaData("trypassword").
		Set().TryPassword(key, count)
	if err != nil {
		return err
	}
	data, err := manager.ToData()
	if err != nil {
		return err
	}
	err = l.store.AddTryPasword(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// GetTryPasword implements Local.
func (l *local) GetTryPasword(ctx context.Context) (int, error) {
	key := l.config.Server().SecretKey(l.config.Client().Version())
	var countTry int
	getCountData, err := l.store.GetTryPasword(ctx)
	if err != nil {
		if err.Error() != "not found" {
			return 0, err
		}
		return 0, err
	}

	counts := getCountData.ToManager()
	countRes, err := counts.AddEncription(l.config.Encryptor()).Get().Data(key)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(countRes, &countTry)
	if err != nil {
		return 0, err
	}
	return countTry, nil
}

// ResetUserData implements Local.
func (l *local) ResetUserData(ctx context.Context) error {
	return l.store.ResetUserData(ctx)
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
func (l *local) GetAccessData(ctx context.Context, masterPsw string) (*core.User, error) {
	data, err := l.store.GetAccessData(ctx)
	if err != nil {
		return nil, err
	}
	// return nil, errors.New("test")
	manager := data.ToManager().AddEncription(l.config.Encryptor())
	// version := l.config.Client().Version()
	// log.Println("===", l.config.Server().SecretKey(version))
	d, err := manager.Get().Data(masterPsw)
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
	return l.store.GetData(ctx, userID, types...)
}

func NewLocal(conf config.Config, store localStore) Local {
	return &local{
		store:  store,
		config: conf,
		auth:   core.NewAuth(conf),
	}
}
