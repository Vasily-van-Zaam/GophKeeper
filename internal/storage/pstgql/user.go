package pstgql

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/jackc/pgx/v5"
)

// GetUserByEmail  returns user by email.
func (s *store) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	var (
		resp  = core.User{}
		err   error
		row   pgx.Row
		query = queryUserByEmail()
	)

	row = s.db.QueryRow(ctx, query, email)

	err = row.Scan(&resp.ID, &resp.Email, &resp.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddUser  insert new user.
func (s *store) AddUser(ctx context.Context, user *core.User) (*core.User, error) {
	var (
		resp  = core.User{}
		err   error
		row   pgx.Row
		query = queryInsertUser()
	)
	args := pgx.NamedArgs{
		"email":      user.Email,
		"privateKey": user.PrivateKey,
	}
	row = s.db.QueryRow(ctx, query, args)
	err = row.Scan(&resp.ID, &resp.Email, &resp.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
