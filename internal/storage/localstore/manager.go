package localstore

import (
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
)

// AddData implements Store.
func (s *store) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	for _, d := range data {
		if d.ID == nil {
			newID := uuid.New()
			d.ID = &newID
		}
	}

	s.data.DataList = append(s.data.DataList, data...)
	err := s.saveToFile()
	if err != nil {
		return nil, err
	}
	return data, nil
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

func (s *store) containsTypes(t string, types ...string) bool {
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
func (s *store) containsSearch(search, body string) bool {
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
