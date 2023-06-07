package server

import (
	context "context"
	"log"
	"net"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ManagerService interface {
	GetData(ctx context.Context, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	SearchData(ctx context.Context, search string, types ...string) ([]*core.ManagerData, error)
}

type UserService interface {
	Login(ctx context.Context, form *core.LoginForm) (*core.AuthToken, error)
	Registration(ctx context.Context, form *core.LoginForm) (*string, error)
	RegistrationAccept(ctx context.Context, form *core.LoginForm) error
}

// Ranner.
type runner interface {
	Run(string) error
	Stop() error
}

type server struct {
	UnimplementedGrpcServer
	config   config.Config
	user     UserService
	service  ManagerService
	listener net.Listener
}

func New(conf config.Config, u UserService, m ManagerService) runner {
	return &server{
		config:  conf,
		user:    u,
		service: m,
	}
}

// Run implements runner.
func (srv *server) Run(addresPort string) error {
	listen, err := net.Listen("tcp", addresPort)
	if err != nil {
		log.Fatal(err)
	}
	srv.listener = listen
	s := grpc.NewServer(grpc.UnaryInterceptor(srv.unaryInterceptor))
	log.Println("Starting grpc server", addresPort)
	RegisterGrpcServer(s, srv)
	return s.Serve(listen)
}

// Stop implements runner.
func (srv *server) Stop() error {
	return srv.listener.Close()
}

// Interceptor implements runner.
func (srv *server) unaryInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var token string
	var clientVersion string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("token")
		client := md.Get("client_version")
		if len(values) > 0 {
			token = values[0]
		}
		if len(client) > 0 {
			clientVersion = client[0]
		}
	}
	log.Println("===", token)
	// if len(token) == 0 {
	// 	return nil, status.Error(codes.Unauthenticated, "missing token")
	// }
	// if token != SecretToken {
	// 	return nil, status.Error(codes.Unauthenticated, "invalid token")
	// }

	// TODO: здесь дописать код который расшифровывает токен
	// получаем id юзера из токена и после добавляем его в metadata user
	key := srv.config.Server().SecretKey(clientVersion)
	md.Set("user", "uniq_user_id")
	md.Set("server_key", key)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return handler(ctx, req)
}
