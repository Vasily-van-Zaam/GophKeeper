// Localstore is used to clients
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package localstore

import (
	"context"
	"encoding/gob"
	"errors"
	"os"
	"sort"
	"strings"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
)

// localstore Implements.
type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)

	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
}

type store struct {
	data     *core.DataGob
	filePath string
	config   config.Config
}

// Close implements Store.
func (*store) Close() error {
	panic("unimplemented")
}

// SearchData implements Store.
func (s *store) SearchData(
	ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	if s.data == nil {
		return nil, errors.New("data is nil")
	}
	res := make([]*core.ManagerData, 0)
	for _, d := range s.data.DataList {
		if d.UserID.String() == userID &&
			s.containsTypes(d.DataType, types...) &&
			s.containsSearch(search, d.MetaData) {
			res = append(res, d)
		}
	}
	return res, nil
}

// AddData implements Store.
func (s *store) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}
	newID := uuid.New()

	for _, d := range data {
		d.ID = &newID
	}
	s.data.DataList = append(s.data.DataList, data...)
	err := s.saveToFile()
	if err != nil {
		return nil, err
	}
	return data, nil
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

func (store) containsTypes(t string, types ...string) bool {
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
func (store) containsSearch(search, body string) bool {
	if search == "" {
		return true
	}
	return strings.Contains(body, search)
}

// GetData implements Store.
func (s *store) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	if s.data == nil {
		return nil, errors.New("data is nil")
	}
	res := make([]*core.ManagerData, 0)
	for _, d := range s.data.DataList {
		if d.UserID.String() == userID && s.containsTypes(d.DataType, types...) {
			res = append(res, d)
		}
	}
	return res, nil
}

// GetUserByEmail implements Store.
func (store) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	panic("unimplemented")
}

// Saving file. Creating a new file, removing old files, renaming new files.
// if the new file was saved without errors, then delete the old one.
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
			data:     &data,
			config:   conf,
			filePath: filePath,
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
