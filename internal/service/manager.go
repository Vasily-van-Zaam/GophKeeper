package service

import (
	"context"
	"errors"
	"log"

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
func (s *service) GetData(ctx context.Context, wihData bool, types ...string) ([]*core.ManagerData, error) {
	var (
		err  error
		resp = make([]*core.ManagerData, 0)
	)

	data, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return nil, errors.New("err metadata")
	}
	userIDS := data.Get("userID")
	userID := ""
	log.Println(userIDS)
	if len(userIDS) == 0 {
		return nil, errors.New("error token metadata")
	}
	userID = userIDS[0]
	if wihData {
		resp, err = s.store.GetData(ctx, userID, types...)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	resp, err = s.store.GetDataInfo(ctx, userID, types...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func New(conf config.Config, store Store, encript core.Encryptor) Service {
	return &service{
		store:     store,
		encriptor: encript,
		config:    conf,
	}
}
