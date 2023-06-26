// Localstore is used to clients
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package localstore

import (
	"context"
	"encoding/gob"
	"errors"
	"os"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

// localstore Implements.
type Store interface {
	Size() int64
	LastSync() string
	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetAccessData(ctx context.Context) (*core.ManagerData, error)
	AddAccessData(ctx context.Context, data *core.ManagerData) error
	// SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
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

// Close implements Store.
func (*store) Close() error {
	return errors.New("not implemented")
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
		info, errStat := dataFile.Stat()
		if errStat != nil {
			return nil, errStat
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
