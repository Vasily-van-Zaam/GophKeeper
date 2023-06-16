package server

import (
	context "context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (srv *server) GetAccess(ctx context.Context, req *GetAccessRequest) (*GetAccessResponse, error) {
	// расшифровываем данные из GetAccessRequest.Email  и ищем по email юзера
	// получаем ответ от сервиса что код отправлен на почту или выдаем ошибку
	var access core.AccessForm
	err := srv.auth.DecryptData(ctx, req.Access, &access)
	if err != nil {
		return nil, err
	}

	token, err := srv.user.GetAccess(ctx, &access)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "not found user")
	}
	return &GetAccessResponse{
		Token: token,
	}, nil
}
func (srv *server) ConfirmAccess(ctx context.Context, req *ConfirmAccessRequest) (*ConfirmAccessResponse, error) {
	// расшифровываем данные из ConfirmAccessRequest.EmailCode
	// дынные будут лежать в виде строки

	var access core.AccessForm
	err := srv.auth.DecryptData(ctx, req.Access, &access)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	user, err := srv.user.ConfirmAccess(ctx, &access)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// зашифровываем данные
	bUser, err := srv.auth.EncryptData(ctx, user)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	return &ConfirmAccessResponse{
		Data: bUser,
	}, nil
}
