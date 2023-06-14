package core

import (
	"context"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	"google.golang.org/grpc/metadata"
)

func Test_authenticated_EncryptDeecryptData(t *testing.T) {
	crypt := cryptor.New()
	type fields struct {
		config config.Config
	}
	type args struct {
		ctx             context.Context
		user            *User
		version         string
		versionOnClient string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "test_version_config_data_0.0.1",
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx:             context.Background(),
				version:         "0.0.1",
				versionOnClient: "0.0.1",
				user: &User{
					Email:      "test@email.com",
					ID:         "123456789",
					PrivateKey: "privatePrivateKey",
				},
			},
			want: &User{
				Email:      "test@email.com",
				ID:         "123456789",
				PrivateKey: "privatePrivateKey",
			},
		},
		{
			name: "test_version_config_data_0.0.0",
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx:             context.Background(),
				version:         "0.0.0",
				versionOnClient: "0.0.0",
				user: &User{
					Email:      "test@email.com",
					ID:         "123456789",
					PrivateKey: "privatePrivateKey",
				},
			},
			want: &User{
				Email:      "test@email.com",
				ID:         "123456789",
				PrivateKey: "privatePrivateKey",
			},
		},
		{
			name: "test_error_config_data",
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.2",
				user: &User{
					Email: "test@email.com",
					ID:    "123456789",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAuth(tt.fields.config)

			md := metadata.New(map[string]string{})
			md.Set("client_version", tt.args.version)
			ctx := metadata.NewIncomingContext(tt.args.ctx, md)

			got, err := a.EncryptData(ctx, tt.args.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("authenticated.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			md.Set("client_version", tt.args.versionOnClient)
			md.Set("token", hex.EncodeToString(got))
			if got == nil && tt.wantErr {
				return
			}
			ctx = metadata.NewIncomingContext(context.Background(), md)
			var user User
			got1 := a.DecryptByContextData(ctx, &user)
			if (got1 != nil) != tt.wantErr {
				t.Errorf("authenticated.GetDataFromToken() error = %v, wantErr %v", got1, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(&user, tt.want) {
				t.Errorf("authenticated.CreateToken() authenticated.GetDataFromToken() = %v, want %v", user, tt.want)
			}
		})
	}
}
