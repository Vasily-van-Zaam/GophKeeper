package pstgql

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
			&info.UserID,
			&info.DataType,
			&info.MetaData,
			&info.Hash,
			&info.UpdatedAt,
			&info.CreatedAt,
		)
		iffc := info.ID.String()
		log.Println(iffc)
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
			&manager.UserID,
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
func (s *store) AddData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	var (
		resp  = 0
		err   error
		query = queryAddData()
	)
	batch := &pgx.Batch{}
	for _, d := range data {
		args := pgx.NamedArgs{
			"id":        d.InfoData.ID,
			"userID":    d.InfoData.UserID,
			"data":      d.Data,
			"dataType":  d.DataType,
			"metaData":  d.MetaData,
			"hash":      d.Hash,
			"updatedAt": d.UpdatedAt,
			"createdAt": d.CreatedAt,
		}
		batch.Queue(query, args)
	}
	results := s.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, d := range data {
		_, err = results.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				log.Printf("data %s already exists", d.MetaData)
				continue
			}

			return 0, fmt.Errorf("unable to insert row: %w", err)
		}
		resp++
	}
	return resp, nil
}

// ChangeData implements Store.
func (s *store) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	var (
		resp  = 0
		err   error
		query = queryChangeData()
	)
	batch := &pgx.Batch{}
	for _, d := range data {
		batch.Queue(
			query,
			d.Data,
			d.DataType,
			d.MetaData,
			d.Hash,
			d.UpdatedAt,
			d.CreatedAt,
		)
	}
	results := s.db.SendBatch(ctx, batch)
	defer results.Close()

	for _, d := range data {
		_, err = results.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				log.Printf("user %s already exists", d.MetaData)
				continue
			}

			return 0, fmt.Errorf("unable to update row: %w", err)
		}
		resp++
	}

	return resp, nil
}
