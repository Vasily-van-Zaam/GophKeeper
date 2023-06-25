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
	GetData(ctx context.Context, withData bool, types ...string) ([]*core.ManagerData, error)
	AddData(ctx context.Context, data ...*core.ManagerData) ([]*core.ManagerData, error)
	ChangeData(ctx context.Context, data ...*core.ManagerData) (int, error)
	SearchData(ctx context.Context, search string, types ...string) ([]*core.ManagerData, error)
}

type UserService interface {
	GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error)
	ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error)
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
	auth     core.Auth
}

func New(conf config.Config, u UserService, m ManagerService) runner {
	return &server{
		config:  conf,
		user:    u,
		service: m,
		auth:    core.NewAuth(conf),
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
	// var token string
	var clientVersion string
	var userID string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		_ = srv.auth.DecryptByContextData(ctx, &userID)
	}

	key := srv.config.Server().SecretKey(clientVersion)
	md.Set("userID", userID)
	md.Set("server_key", key)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return handler(ctx, req)
}
