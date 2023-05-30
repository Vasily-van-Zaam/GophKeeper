// The Service package is main business logic
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type Store interface {
	GetData(ctx context.Context, userID string, types ...string) ([]core.Manager, error)
	// GetDataById(ctx context.Context, id string) ([]core.Manager, error)
	// GetDataByMeta(ctx context.Context, userID string, types ...string) ([]core.Manager, error)
	AddData(ctx context.Context, data core.Manager) error
	ChangeData(ctx context.Context, data ...core.Manager) (int, error)
}
type Service interface {
	GetData(ctx context.Context, types ...core.DataType) ([]core.Manager, error)
	AddData(ctx context.Context, data core.Manager) error
	ChangeData(ctx context.Context, data ...core.Manager) (int, error)
}

type service struct {
	store     Store
	encriptor core.Encryptor
}

// AddData implements Service.
func (*service) AddData(ctx context.Context, data core.Manager) error {
	panic("unimplemented")
}

// ChangeData implements Service.
func (*service) ChangeData(ctx context.Context, data ...core.Manager) (int, error) {
	panic("unimplemented")
}

// GetData implements Service.
func (*service) GetData(ctx context.Context, types ...core.DataType) ([]core.Manager, error) {
	panic("unimplemented")
}

func New(store Store, encript core.Encryptor) Service {
	return &service{
		store:     store,
		encriptor: encript,
	}
}
