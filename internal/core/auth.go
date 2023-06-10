package core

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"google.golang.org/grpc/metadata"
)

// Auth functions to check authentication.
type Authenticated interface {
	CreateToken(ctx context.Context, user *User) ([]byte, error)
	GetDataFromToken(ctx context.Context) (*User, error)
}
type authenticated struct {
	config config.Config
}

// GetDataFromToken implements Authenticated.
func (a *authenticated) GetDataFromToken(ctx context.Context) (*User, error) {
	const timDelta = 30
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("err metadata")
	}
	version := data.Get("client_version")
	if len(version) == 0 {
		return nil, errors.New("err metadata client version")
	}
	key := a.config.Server().SecretKey(version[0])
	if key == "" {
		return nil, errors.New("version not found")
	}
	tokens := data.Get("token")
	if len(tokens) == 0 {
		return nil, errors.New("unauthenticated 1")
	}
	token := tokens[0]
	if token == "" {
		return nil, errors.New("unauthenticated 2")
	}
	nowPlus30 := time.Now().UTC().Add(time.Second * timDelta).Format("2006-01-02T15:04")
	nowMinus30 := time.Now().UTC().Add(-time.Second * timDelta).Format("2006-01-02T15:04")
	bToken, err := hex.DecodeString(token)
	if err != nil {
		return nil, errors.New("unauthenticated 3")
	}

	manager := NewManagerFromData(&ManagerData{
		Data: bToken,
	}).AddEncription(a.config.Encryptor())
	dec, err := manager.Get().Data(key + nowPlus30)
	if err != nil {
		dec, err = manager.Get().Data(key + nowMinus30)
		if err != nil {
			return nil, err
		}
	}
	var user User
	err = json.Unmarshal(dec, &user)
	if err != nil {
		return nil, err
	}

	log.Println(dec)
	log.Print(data)
	return &User{}, nil
}

// CreateToken implements Authenticated.
func (a *authenticated) CreateToken(ctx context.Context, user *User) ([]byte, error) {
	data, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("err metadata")
	}
	vesion := data.Get("client_version")
	if len(vesion) == 0 {
		return nil, errors.New("err metadata client version")
	}
	secretKey := a.config.Server().SecretKey(vesion[0])
	if secretKey == "" {
		return nil, errors.New("version not found")
	}

	manager := NewManager().AddEncription(a.config.Encryptor())

	userEnc := &User{
		Email:     user.Email,
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Hash:      user.Hash,
	}

	userB, _ := json.Marshal(userEnc)
	now := time.Now().UTC().Format("2006-01-02T15:04")
	err := manager.Set().Data(secretKey+now, userB)
	if err != nil {
		return nil, err
	}
	return manager.Get().EncryptData(), nil
}

func NewAuth(conf config.Config) Authenticated {
	return &authenticated{
		config: conf,
	}
}
