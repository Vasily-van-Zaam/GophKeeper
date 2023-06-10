package service

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	"google.golang.org/grpc/metadata"
)

type mockStore struct {
	Data []core.Manager
}

func Test_service_GetData(t *testing.T) {
	type args struct {
		userID string
		types  []string
	}
	tests := []struct {
		name string
		s    *service
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s := New
			// s.GetData(tt.args.userID, tt.args.types...)
			var sh core.PasswordForm
			p := &core.PasswordForm{
				Login: "ok",
			}
			b, _ := json.Marshal(p)

			json.Unmarshal(b, &sh)

			log.Println(sh)
		})
	}
}

func Test_service_handlerAuth(t *testing.T) {
	crypt := cryptor.New()
	type fields struct {
		store     Store
		encriptor core.Encryptor
		config    config.Config
	}
	type args struct {
		ctx  context.Context
		user *core.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.AuthToken
		wantErr bool
	}{
		{
			fields: fields{
				config: config.New(logger.New(), crypt),
			},
			args: args{
				ctx: context.Background(),
				user: &core.User{
					Email: "email@example.com",
					Hash:  "1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store:     tt.fields.store,
				encriptor: tt.fields.encriptor,
				config:    tt.fields.config,
			}

			md := metadata.New(map[string]string{})
			md.Set("client_version", "0.0.1")

			ctx := metadata.NewIncomingContext(tt.args.ctx, md)

			got, err := s.handlerAuth(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.handlerAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.handlerAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
