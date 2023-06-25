package localstore

import (
	"context"
	"errors"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
)

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
