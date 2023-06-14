package server

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	"google.golang.org/grpc/metadata"
)

func (*userServiceMock) GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error) {
	if form.Email == "email@example.com" {
		return nil, nil
	}
	if form.Email == "" {
		return nil, fmt.Errorf("email is error")
	}
	return nil, fmt.Errorf("email is empty")
}

func (*userServiceMock) ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error) {
	if form.Email == "email@example.com" && form.Code == "123456789" {
		return &core.User{
			Email:      "email@example.com",
			PrivateKey: "private_sevret_user_key",
		}, nil
	}

	return nil, fmt.Errorf("mail or code is not correct")
}

type userServiceMock struct {
}

func Test_server_GetAccess(t *testing.T) {
	type fields struct {
		config config.Config
		user   UserService
		auth   core.Auth
	}
	type args struct {
		ctx     context.Context
		version string
		access  core.AccessForm
	}

	conf := config.New(logger.New(), cryptor.New())
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAccessResponse
		wantErr bool
	}{
		{
			name: "test_get_access_response_ok",
			fields: fields{
				config: conf,
				user:   &userServiceMock{},
				auth:   core.NewAuth(conf),
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.0",
				access: core.AccessForm{
					Email: "email@example.com",
				},
			},
			want: &GetAccessResponse{
				Token: []byte("code sent to email"),
			},
		},
		{
			name: "test_get_access_response_eror_version_not_founed",
			fields: fields{
				config: conf,
				user:   &userServiceMock{},
				auth:   core.NewAuth(conf),
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.2",
				access: core.AccessForm{
					Email: "email@example.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &server{
				config: tt.fields.config,
				user:   tt.fields.user,
				auth:   tt.fields.auth,
			}

			md := metadata.New(map[string]string{core.CtxVersionClientKey: tt.args.version})
			ctx := metadata.NewIncomingContext(tt.args.ctx, md)
			access, err := tt.fields.auth.EncryptData(ctx, tt.args.access)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if access == nil && tt.wantErr {
				return
			}
			got, err := srv.GetAccess(ctx, &GetAccessRequest{
				Access: access,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ConfirmAccess(t *testing.T) {
	conf := config.New(logger.New(), cryptor.New())
	type fields struct {
		config config.Config
		user   UserService
		auth   core.Auth
	}
	type args struct {
		ctx     context.Context
		version string
		access  *core.AccessForm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name: "test_confirm_access_ok",
			fields: fields{
				config: conf,
				auth:   core.NewAuth(conf),
				user:   &userServiceMock{},
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.0",
				access: &core.AccessForm{
					Email: "email@example.com",
					Code:  "123456789",
				},
			},
			want: &core.User{
				Email:      "email@example.com",
				PrivateKey: "private_sevret_user_key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &server{
				config: tt.fields.config,
				user:   tt.fields.user,
				auth:   tt.fields.auth,
			}

			md := metadata.New(map[string]string{core.CtxVersionClientKey: tt.args.version})
			ctx := metadata.NewIncomingContext(tt.args.ctx, md)
			access, err := tt.fields.auth.EncryptData(ctx, tt.args.access)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if access == nil && tt.wantErr {
				return
			}
			got, err := srv.ConfirmAccess(ctx, &ConfirmAccessRequest{
				Access: access,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("server.ConfirmAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var user core.User
			err = srv.auth.DecryptData(ctx, got.Data, &user)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(&user, tt.want) {
				t.Errorf("server.ConfirmAccess() = %v, want %v", user, tt.want)
			}
		})
	}
}

func TestJson(t *testing.T) {
	var str string

	err := json.Unmarshal([]byte("{}"), &str)
	log.Println(err, str)
}
