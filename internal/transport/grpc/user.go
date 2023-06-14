package server

import (
	context "context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
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
		return nil, err
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
		return nil, err
	}

	user, err := srv.user.ConfirmAccess(ctx, &access)

	if err != nil {
		return nil, err
	}

	// зашифровываем данные
	bUser, err := srv.auth.EncryptData(ctx, user)
	if err != nil {
		return nil, err
	}
	return &ConfirmAccessResponse{
		Data: bUser,
	}, nil
}
