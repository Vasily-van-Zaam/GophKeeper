package server

import (
	context "context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

func (srv *server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var (
		resp *LoginResponse
	)
	res, err := srv.user.Login(ctx, &core.LoginForm{
		Email:   req.Email,
		Pasword: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp = &LoginResponse{
		Access:  res.Access,
		Refresh: res.Refresh,
		UserKey: res.UserKey,
	}

	return resp, nil
}
