package server

import (
	context "context"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
)

func (srv *server) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	var (
		resp = &GetDataResponse{}
	)

	res, err := srv.service.GetData(ctx, req.WithData, req.DataTypes...)
	if err != nil {
		return nil, err
	}
	for _, d := range res {
		uID := ""
		if d.UserID != nil {
			uID = d.UserID.String()
		}
		resp.List = append(resp.List, &ManagerData{
			ID:        d.ID.String(),
			DataType:  d.DataType,
			MetaData:  d.MetaData,
			UserID:    uID,
			Local:     d.Local,
			Hash:      d.Hash,
			CreatedAt: d.CreatedAt.Format(time.RFC3339),
			UpdatedAt: d.UpdatedAt.Format(time.RFC3339),
			Data:      d.Data,
		})
	}

	return resp, nil
}

func (srv *server) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return &PingResponse{
		Ok: true,
	}, nil
}

func (srv *server) ChangeData(ctx context.Context, req *ChangeDataRequest) (*ChangeDataResponse, error) {
	return nil, nil
}
func (srv *server) AddData(ctx context.Context, req *AddDataRequest) (*AddDataResponse, error) {
	var (
		err   error
		added int
		data  = make([]*core.ManagerData, len(req.List))
	)
	for i, v := range req.List {
		userID, _ := uuid.Parse(v.UserID)
		id, _ := uuid.Parse(v.ID)
		createdAt, _ := time.Parse(time.RFC3339, v.CreatedAt)
		updateAt, _ := time.Parse(time.RFC3339, v.UpdatedAt)

		data[i] = &core.ManagerData{
			InfoData: core.InfoData{
				ID:        &id,
				UserID:    &userID,
				MetaData:  v.MetaData,
				DataType:  v.DataType,
				Hash:      v.Hash,
				CreatedAt: &createdAt,
				UpdatedAt: &updateAt,
			},
			Data: v.Data,
		}
	}
	added, err = srv.service.AddData(ctx, data...)
	if err != nil {
		return nil, err
	}
	return &AddDataResponse{
		Added: int32(added),
		Error: "",
	}, nil
}
