package server

import (
	context "context"
	"log"
	"net"
	reflect "reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type userMockService struct {
}

func (s *userMockService) Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error) {
	data, ok := metadata.FromIncomingContext(ctx)
	log.Println(data, ok)
	return &core.AuthToken{
		Access:  []byte("access"),
		Refresh: []byte("refresh"),
		UserKey: []byte("user_key"),
	}, nil // errors.New("error login")
}
func (s *userMockService) Registration(ctx context.Context, form *core.LoginForm) (*string, error) {
	return nil, nil
}
func (s *userMockService) RegistrationAccept(ctx context.Context, form *core.LoginForm) error {
	return nil
}

const addresPort = "localhost:3200"

func Test_server_Login(t *testing.T) {
	type fields struct {
		UnimplementedGrpcServer UnimplementedGrpcServer
		config                  config.Config
		user                    UserService
		service                 ManagerService
		listener                net.Listener
	}

	f := fields{
		UnimplementedGrpcServer: UnimplementedGrpcServer{},
		config:                  config.New(logger.New()),
		user:                    &userMockService{},
		service:                 nil,
		listener:                nil,
	}
	type args struct {
		ctx context.Context
		req *LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginResponse
		wantErr bool
	}{
		{
			args: args{
				req: &LoginRequest{
					Email:    "mail@mail.ru",
					Password: "pasword",
				},
			},
			want: &LoginResponse{
				Access:  []byte("access"),
				Refresh: []byte("refresh"),
				UserKey: []byte("user_key"),
			},
		},
	}

	srv := New(f.config, f.user, f.service)
	log.Println(srv)
	// go srv.Run(addresPort)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := grpc.Dial(addresPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			c := NewGrpcClient(conn)
			ctx := context.Background()
			md := metadata.New(map[string]string{"client_version": "0.1.1"})
			ctx = metadata.NewOutgoingContext(context.Background(), md)
			got, err := c.Login(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Access, tt.want.Access) {
				t.Errorf("server.Login() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Refresh, tt.want.Refresh) {
				t.Errorf("server.Login() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.UserKey, tt.want.UserKey) {
				t.Errorf("server.Login() = %v, want %v", got, tt.want)
			}
			log.Printf("stop server")
		})
	}
	// defer func() {
	// 	err := srv.Stop()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }()
}
