package server

import (
	context "context"
	"time"
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
	return nil, nil
}
