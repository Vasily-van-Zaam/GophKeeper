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
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
)

// localstore Implements.
type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	Size() int64
	LastSync() string
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context) (*core.ManagerData, error)
	AddAccessData(ctx context.Context, data *core.ManagerData) error
	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	Close() error
	ResetUserData(ctx context.Context, types ...core.DataType) error
	AddTryPasword(ctx context.Context, data *core.ManagerData) error
	GetTryPasword(ctx context.Context) (*core.ManagerData, error)
}

type store struct {
	data     *core.DataGob
	filePath string
	config   config.Config
	size     int64
	lastSync *time.Time
}

// SLastSync implements Store.
func (s *store) LastSync() string {
	if s.lastSync == nil {
		return "never"
	}
	return s.lastSync.Local().Format("2006-01-02 15:04")
}

// Size implements Store.
func (s store) Size() int64 {
	return s.size
}

// AddAccessData implements Store.
func (s *store) AddAccessData(ctx context.Context, data *core.ManagerData) error {
	if data == nil {
		return errors.New("data is nil")
	}

	newID := uuid.New()
	exists := false
	for i, d := range s.data.DataList {
		if d.DataType == string(core.DataTypeUser) {
			d = data
			exists = true
			s.data.DataList[i] = s.data.DataList[len(s.data.DataList)-1]
			s.data.DataList = s.data.DataList[:len(s.data.DataList)-1]
			break
		}
	}
	for _, d := range s.data.DataList {
		if d.DataType == string(core.DataTypeTryEnterPassword) {
			d = data
			exists = true
			break
		}
	}
	if !exists {
		data.ID = &newID
		s.data.DataList = append(s.data.DataList, data)
	}

	err := s.saveToFile()
	if err != nil {
		return err
	}

	return nil
}

// AddTryEnterPasword implements Store.
func (s *store) AddTryPasword(ctx context.Context, data *core.ManagerData) error {
	if s.data == nil {
		return errors.New("data is nil")
	}
	exists := false
	for i, d := range s.data.DataList {
		if d.DataType == string(core.DataTypeTryEnterPassword) {
			s.data.DataList[i] = data
			exists = true

			err := s.saveToFile()
			if err != nil {
				return err
			}
			break
		}
	}
	if !exists {
		s.data.DataList = append(s.data.DataList, data)
		err := s.saveToFile()
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTryEnterPasword implements Store.
func (s *store) GetTryPasword(ctx context.Context) (*core.ManagerData, error) {
	if s.data == nil {
		return nil, errors.New("data is nil")
	}
	for _, d := range s.data.DataList {
		if d.DataType == string(core.DataTypeTryEnterPassword) {
			return d, nil
		}
	}
	return nil, errors.New("not found")
}

// ResetUserData implements Store.
func (s *store) ResetUserData(ctx context.Context, types ...core.DataType) error {
	if s.data == nil {
		return errors.New("data is nil")
	}
	deleteIndexes := make([]int, 0)
	check := func(i int) bool {
		for _, n := range deleteIndexes {
			if i == n {
				return true
			}
		}
		return false
	}
	newList := make([]*core.ManagerData, 0)

	for _, t := range types {
		for i, d := range s.data.DataList {
			if string(t) == d.DataType {
				deleteIndexes = append(deleteIndexes, i)
			}
		}
	}
	for i, d := range s.data.DataList {
		if !check(i) {
			newList = append(newList, d)
		}
	}
	s.data.DataList = newList
	err := s.saveToFile()
	if err != nil {
		return err
	}
	return nil
}

// GetAccessData implements Store.
func (s *store) GetAccessData(ctx context.Context) (*core.ManagerData, error) {
	if s.data == nil {
		return nil, errors.New("data is nil")
	}

	for _, d := range s.data.DataList {
		if d.DataType == string(core.DataTypeUser) {
			return d, nil
		}
	}
	return nil, errors.New("not found")
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

	for _, d := range data {
		newID := uuid.New()
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
func (a *store) AddUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// ChangeData implements Store.
func (s *store) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	changed := 0
	for i, d := range s.data.DataList {
		for _, e := range data {
			if d.InfoData.ID == e.InfoData.ID {
				s.data.DataList[i] = e
				changed++
			}
		}
	}
	err := s.saveToFile()
	if err != nil {
		return 0, err
	}
	return changed, nil
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
		if d.UserID == nil {
			continue
			// return nil, errors.New("userID is nil")
		}
		// res = append(res, d)
		id := d.UserID.String()
		if id == userID && s.containsTypes(d.DataType, types...) {
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

	data := s.data
	enc := gob.NewEncoder(dataFile)
	err = enc.Encode(&data)
	if err != nil {
		return err
	}
	info, err := dataFile.Stat()

	if err != nil {
		return err
	}
	s.size = info.Size()
	err = os.Remove(s.filePath)
	if err != nil {
		s.config.Logger().Error(err)
	}
	dataFile.Close()
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
		info, err := dataFile.Stat()
		if err != nil {
			return nil, err
		}
		size := info.Size()
		return &store{
			data:     &data,
			config:   conf,
			filePath: filePath,
			size:     size,
		}, nil
	}
	defer dataFile.Close()
	var data *core.DataGob
	dec := gob.NewDecoder(dataFile)
	err = dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	info, err := dataFile.Stat()
	if err != nil {
		return nil, err
	}
	size := info.Size()
	return &store{
		data:     data,
		filePath: filePath,
		config:   conf,
		size:     size,
	}, nil
}
