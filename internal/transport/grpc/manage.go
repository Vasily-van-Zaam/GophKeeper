package server

import (
	context "context"
	"time"
)

func (srv *server) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	var (
		resp = &GetDataResponse{}
	)
	res, err := srv.service.GetData(ctx, req.DataTypes...)
	if err != nil {
		return nil, err
	}
	for _, d := range res {
		resp.List = append(resp.List, &ManagerData{
			ID:        d.ID.String(),
			DataType:  d.DataType,
			MetaData:  d.MetaData,
			UserID:    d.UserID.String(),
			Local:     d.Local,
			Hash:      d.Hash,
			CreatedAt: d.CreatedAt.Format(time.RFC3339),
			UpdatedAt: d.UpdatedAt.Format(time.RFC3339),
			Data:      d.Data,
		})
	}

	return resp, nil
}
