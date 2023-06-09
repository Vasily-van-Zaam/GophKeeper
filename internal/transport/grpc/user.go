package server

import (
	context "context"
	"strings"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
		if strings.Contains(err.Error(), "not found") {
			return nil, status.Errorf(codes.PermissionDenied, "incorrect login or password")
		}
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	resp = &LoginResponse{
		Access:  res.Access,
		Refresh: res.Refresh,
		UserKey: res.UserKey,
	}

	return resp, nil
}
