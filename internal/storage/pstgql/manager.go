package pstgql

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/jackc/pgx/v5"
)

// GetData - returns information about manager data user.
func (s *store) GetDataInfo(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	var (
		resp  = make([]*core.ManagerData, 0)
		err   error
		rows  pgx.Rows
		query = queryGetUserData(false, types...)
	)

	rows, err = s.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	s.config.Logger().Debug(query)
	for rows.Next() {
		info := core.ManagerData{}

		err = rows.Scan(
			&info.ID,
			&info.DataType,
			&info.MetaData,
			&info.Hash,
			&info.UpdatedAt,
			&info.CreatedAt,
		)
		if err != nil {
			s.config.Logger().Error(err)
		}
		resp = append(resp, &info)
	}

	return resp, nil
}

// GetData returns all duser managet dada.
func (s *store) GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error) {
	var (
		resp  = make([]*core.ManagerData, 0)
		err   error
		rows  pgx.Rows
		query = queryGetUserData(true, types...)
	)

	rows, err = s.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	s.config.Logger().Debug(query)
	for rows.Next() {
		manager := core.ManagerData{}

		err = rows.Scan(
			&manager.ID,
			&manager.Data,
			&manager.DataType,
			&manager.MetaData,
			&manager.Hash,
			&manager.UpdatedAt,
			&manager.CreatedAt,
		)
		if err != nil {
			s.config.Logger().Error(err)
		}
		resp = append(resp, &manager)
	}

	return resp, nil
}

// AddData implements Store.
func (s *store) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	// var (
	// 	resp  = make([]*core.ManagerData, 0)
	// 	err   error
	// 	rows  pgx.Rows
	// 	query = queryGetUserData(true, types...)
	// )

	return nil, nil
}

// ChangeData implements Store.
func (*store) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}
