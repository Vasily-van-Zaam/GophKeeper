// Store on server Postgres
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package pstgql

import (
	"context"
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Implements store functionality.
type Store interface {
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)
	AddUser(ctx context.Context, user *core.User) (*core.User, error)
	ChangeUser(ctx context.Context, user *core.User) (*core.User, error)
	GetSecretToken(ctx context.Context, userID string) ([]byte, error)

	GetData(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)
	GetDataInfo(ctx context.Context, userID string, types ...string) ([]*core.ManagerData, error)

	AddData(ctx context.Context, data ...*core.ManagerData) (int, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)

	SearchData(ctx context.Context, search, userID string, types ...string) ([]*core.ManagerData, error)
}

type store struct {
	config config.Config
	db     *pgxpool.Pool
}

// ChangeUser implements Store.
func (s *store) ChangeUser(ctx context.Context, user *core.User) (*core.User, error) {
	panic("unimplemented")
}

// GetSecretToken implements Store.
func (*store) GetSecretToken(ctx context.Context, userID string) ([]byte, error) {
	panic("unimplemented")
}

// SearchData implements Store.
func (*store) SearchData(ctx context.Context, search string, userID string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

func New(conf config.Config) (Store, error) {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(conf.Server().DataBaseDNS())
	if err != nil {
		panic(err)
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		return nil
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		panic(err)
	}

	_, errExecUser := db.Exec(ctx, `--sql
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
	CREATE TABLE IF NOT EXISTS users(
		id uuid UNIQUE DEFAULT uuid_generate_v1() PRIMARY KEY,
		private_key character varying  NOT NULL,
		email character varying UNIQUE,
		email_confirmed boolean default false
	);`)
	_, errExec := db.Exec(ctx, `--sql
	CREATE TABLE IF NOT EXISTS manager_data(
		id uuid PRIMARY KEY UNIQUE,
		data bytea,
		data_type character varying,
		meta_data character varying,
		hash character varying,
		updated_at timestamp with time zone,
		created_at timestamp with time zone,
		user_id uuid REFERENCES users (id)
	);`)

	if errExecUser != nil {
		log.Println("errExecUser: ", errExecUser.Error())
	}

	if errExec != nil {
		log.Println("errExec: ", errExec.Error())
	}

	return &store{
		config: conf,
		db:     db,
	}, nil
}
