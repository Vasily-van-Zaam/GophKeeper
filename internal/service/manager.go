package service

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type service struct {
	store     Store
	encriptor core.Encryptor
}

// AddData implements Service.
func (*service) AddData(ctx context.Context, data *core.ManagerData) error {
	panic("unimplemented")
}

// ChangeData implements Service.
func (*service) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// GetData implements Service.
func (*service) GetData(ctx context.Context, types ...core.DataType) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

func New(store Store, encript core.Encryptor) Service {
	return &service{
		store:     store,
		encriptor: encript,
	}
}
