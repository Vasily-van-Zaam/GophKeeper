package localstore

import (
	"context"
	"encoding/gob"
	"errors"
	"os"
	"sort"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)

	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	SearchDataByMeta(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data *core.ManagerData) (*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
}

type store struct {
	data     *core.DataGob
	filePath string
	config   config.Config
}

// SearchDataByMeta implements Store.
func (*store) SearchDataByMeta(ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// AddData implements Store.
func (*store) AddData(ctx context.Context, data *core.ManagerData) (*core.ManagerData, error) {
	panic("unimplemented")
}

// AddUser implements Store.
func (store) AddUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// ChangeData implements Store.
func (store) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// ChangeUser implements Store.
func (store) ChangeUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// GetData implements Store.
func (s *store) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	if s.data == nil {
		return nil, errors.New("data is nil")
	}
	getTypes := func(t string) bool {
		sort.Strings(types)
		i := sort.SearchStrings(types, t)
		if i < len(types) && types[i] == t {
			return true
		}
		if len(types) == 0 {
			return true
		}
		return false
	}
	res := make([]*core.ManagerData, 0)
	for _, d := range s.data.DataList {
		if d.UserID.String() == userID && getTypes(d.DataType) {
			res = append(res, d)
		}
	}
	return res, nil
}

// GetUserByEmail implements Store.
func (store) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	panic("unimplemented")
}

func (s *store) saveToFile() error {
	dataFile, err := os.Create("." + s.filePath)
	if err != nil {
		return err
	}
	defer dataFile.Close()
	data := s.data
	enc := gob.NewEncoder(dataFile)
	err = enc.Encode(&data)
	if err != nil {
		return err
	}
	err = os.Remove(s.filePath)
	if err != nil {
		s.config.Logger().Error(err)
	}
	err = os.Rename("."+s.filePath, s.filePath)
	if err != nil {
		return err
	}

	return nil
}

// Create new store.
func New(conf config.Config) (Store, error) {
	filePath := conf.Client().FilePath()

	dataFile, err := os.Open(filePath)
	if err != nil {
		dataFile, err = os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer dataFile.Close()
		data := core.DataGob{}
		enc := gob.NewEncoder(dataFile)
		err = enc.Encode(&data)
		if err != nil {
			return nil, err
		}
		return &store{
			data: &data,
		}, nil
	}
	defer dataFile.Close()
	var data *core.DataGob
	dec := gob.NewDecoder(dataFile)
	err = dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &store{
		data:     data,
		filePath: filePath,
		config:   conf,
	}, nil
}
