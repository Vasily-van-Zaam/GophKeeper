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

func Test_authenticated_CreateToken(t *testing.T) {
	crypt := cryptor.New()
	type fields struct {
		config config.Config
	}
	type args struct {
		ctx     context.Context
		user    *User
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name: "test_version_config_data_0.0.1",
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.1",
				user: &User{
					Email: "test@email.com",
					ID:    "123456789",
				},
			},
			wantLen: 195,
		},
		{
			name: "test_version_config_data_0.0.0",
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx:     context.Background(),
				version: "0.0.0",
				user: &User{
					Email: "test@email.com",
					ID:    "123456789",
				},
			},
			wantLen: 195,
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

			got, err := a.CreateToken(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("authenticated.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("authenticated.CreateToken() = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}

func Test_authenticated_GetDataFromToken(t *testing.T) {
	type fields struct {
		config config.Config
	}
	type args struct {
		ctx   context.Context
		token []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			fields: fields{
				config: config.New(logger.New(), cryptor.New()),
			},
			args: args{
				ctx:   context.Background(),
				token: []byte("token"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAuth(tt.fields.config)
			md := metadata.New(map[string]string{})
			md.Set("client_version", "0.0.0")
			md.Set("token", hex.EncodeToString(tt.args.token))

			ctx := metadata.NewIncomingContext(tt.args.ctx, md)
			got, err := a.GetDataFromToken(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("authenticated.GetDataFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authenticated.GetDataFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
