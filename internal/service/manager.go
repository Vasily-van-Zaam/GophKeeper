package service

import (
	"context"
	"errors"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"google.golang.org/grpc/metadata"
)

type service struct {
	store     Store
	encriptor core.Encryptor
	config    config.Config
}

// SearchData implements Service.
func (*service) SearchData(ctx context.Context, search string, types ...string) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// AddData implements Service.
func (*service) AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error) {
	panic("unimplemented")
}

// ChangeData implements Service.
func (*service) ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error) {
	panic("unimplemented")
}

// GetData implements Service.
func (s *service) GetData(ctx context.Context, types ...string) ([]*core.ManagerData, error) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("err metadata")
	}
	data.Get("user")

	// s.store.GetData(ctx, ...types)
	return nil, nil
}

func New(conf config.Config, store Store, encript core.Encryptor) Service {
	return &service{
		store:     store,
		encriptor: encript,
		config:    conf,
	}
}
