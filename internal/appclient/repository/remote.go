package repository

import (
	"context"
	"encoding/hex"
	"log"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	server "github.com/Vasily-van-Zaam/GophKeeper.git/internal/transport/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type remoteStore interface {
	GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error)
	ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error)
	GetData(ctx context.Context, types ...string) ([]*core.ManagerData, error)
	Ping(ctx context.Context) (bool, error)
	Close() error
	// SyncData(ctx context.Context) ()
}

type remote struct {
	config config.Config
	conn   *grpc.ClientConn
	client server.GrpcClient
	auth   core.Auth
}

// Close implements Remote.
func (r *remote) Close() error {
	return r.conn.Close()
}

// ConfirmAccess implements Remote.
func (r *remote) ConfirmAccess(ctx context.Context, form *core.AccessForm) (*core.User, error) {
	md := metadata.New(map[string]string{
		core.CtxVersionClientKey: r.config.Client().Version(),
		core.CtxAcceptToken:      hex.EncodeToString(form.Token),
	})

	ctx = metadata.NewIncomingContext(ctx, md)
	access, err := r.auth.EncryptData(ctx, form)
	if err != nil {
		// log.Print(ctx)
		return nil, err
	}
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := r.client.ConfirmAccess(ctx, &server.ConfirmAccessRequest{
		Access: access,
	})
	if err != nil {
		return nil, err
	}
	var user *core.User
	err = r.auth.DecryptData(ctx, resp.Data, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAccess implements Remote.
func (r *remote) GetAccess(ctx context.Context, form *core.AccessForm) ([]byte, error) {
	md := metadata.New(map[string]string{core.CtxVersionClientKey: r.config.Client().Version()})

	ctx = metadata.NewIncomingContext(ctx, md)
	access, err := r.auth.EncryptData(ctx, form)
	if err != nil {
		// log.Print(ctx)
		return nil, err
	}
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp, err := r.client.GetAccess(ctx, &server.GetAccessRequest{
		Access: access,
	})
	if err != nil {
		return nil, err
	}
	return resp.Token, nil
}

// GetData implements Remote.
func (r *remote) GetData(ctx context.Context, types ...string) ([]*core.ManagerData, error) {
	resp, err := r.client.GetData(ctx, &server.GetDataRequest{
		DataTypes: types,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*core.ManagerData, len(resp.List))
	for i, d := range resp.List {
		uID, _ := uuid.Parse(d.UserID)
		cDate, _ := time.Parse(time.RFC3339, d.CreatedAt)
		uDate, _ := time.Parse(time.RFC3339, d.UpdatedAt)
		data[i] = &core.ManagerData{
			InfoData: core.InfoData{
				MetaData:  d.MetaData,
				UserID:    &uID,
				Local:     d.Local,
				DataType:  d.DataType,
				Hash:      d.Hash,
				CreatedAt: &cDate,
				UpdatedAt: &uDate,
			},
			Data: d.Data,
		}
	}

	return data, nil
}

// Ping implements Remote.
func (r *remote) Ping(ctx context.Context) (bool, error) {
	resp, err := r.client.Ping(ctx, nil)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func NewRemote(conf config.Config) remoteStore {
	conn, err := grpc.Dial(conf.Client().SrvAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("connection to prod err", err)
	}
	client := server.NewGrpcClient(conn)
	_, err = client.Ping(context.Background(), nil)

	if err != nil {
		conn, err = grpc.Dial(conf.Client().SrvAddressProd(), grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Println("connection to prod err", err)
		}
		client = server.NewGrpcClient(conn)
	}

	return &remote{
		config: conf,
		conn:   conn,
		client: client,
		auth:   core.NewAuth(conf),
	}
}
