package core

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"google.golang.org/grpc/metadata"
)

// Auth functions to check authentication.
type Auth interface {
	DecryptByContextData(ctx context.Context, d any) error
	DecryptData(ctx context.Context, enc []byte, dec any) error
	EncryptData(ctx context.Context, d any) ([]byte, error)
}
type auth struct {
	config config.Config
}

// DecryptData implements Authenticated.
func (a *auth) DecryptData(ctx context.Context, enc []byte, dec any) error {
	const timDelta = 30
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("err metadata")
	}
	version := data.Get(CtxVersionClientKey)
	if len(version) == 0 {
		return errors.New("err metadata client version")
	}
	key := a.config.Server().SecretKey(version[0])
	if key == "" {
		return errors.New("version not found")
	}

	nowPlus30 := time.Now().UTC().Add(time.Second * timDelta).Format("2006-01-02T15:04")
	nowMinus30 := time.Now().UTC().Add(-time.Second * timDelta).Format("2006-01-02T15:04")
	// log.Println(nowPlus30, nowMinus30)
	manager := NewManagerFromData(&ManagerData{
		Data: enc,
	}).AddEncription(a.config.Encryptor())
	decryption, err := manager.Get().Data(key + nowMinus30)
	if err != nil {
		decryption, err = manager.Get().Data(key + nowPlus30)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(decryption, dec)
	if err != nil {
		return err
	}

	return nil
}

// DecryptData implements Authenticated.
func (a *auth) DecryptByContextData(ctx context.Context, d any) error {
	const timDelta = 30
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("err metadata")
	}
	version := data.Get("client_version")
	if len(version) == 0 {
		return errors.New("err metadata client version")
	}
	key := a.config.Server().SecretKey(version[0])
	if key == "" {
		return errors.New("version not found")
	}
	tokens := data.Get(CtxTokenKey)
	if len(tokens) == 0 {
		return errors.New("unauthenticated 1")
	}
	token := tokens[0]
	if token == "" {
		return errors.New("unauthenticated 2")
	}
	nowPlus30 := time.Now().UTC().Add(time.Second * timDelta).Format("2006-01-02T15:04")
	nowMinus30 := time.Now().UTC().Add(-time.Second * timDelta).Format("2006-01-02T15:04")
	// log.Println(nowPlus30, nowMinus30)
	bToken, err := hex.DecodeString(token)
	if err != nil {
		return errors.New("unauthenticated 3")
	}

	manager := NewManagerFromData(&ManagerData{
		Data: bToken,
	}).AddEncription(a.config.Encryptor())
	dec, err := manager.Get().Data(key + nowMinus30)
	if err != nil {
		dec, err = manager.Get().Data(key + nowPlus30)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(dec, d)
	if err != nil {
		return err
	}

	return nil
}

// EncryptDatas implements Authenticated.
func (a *auth) EncryptData(ctx context.Context, d any) ([]byte, error) {
	data, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("err metadata")
	}
	vesion := data.Get(CtxVersionClientKey)
	if len(vesion) == 0 {
		return nil, errors.New("err metadata client version")
	}
	secretKey := a.config.Server().SecretKey(vesion[0])
	if secretKey == "" {
		return nil, errors.New("version not found")
	}

	manager := NewManager(nil).AddEncription(a.config.Encryptor())

	userB, _ := json.Marshal(d)
	now := time.Now().UTC().Format("2006-01-02T15:04")
	// log.Println(now)
	err := manager.Set().Data(secretKey+now, userB)
	if err != nil {
		return nil, err
	}
	return manager.Get().EncryptData(), nil
}

func NewAuth(conf config.Config) Auth {
	return &auth{
		config: conf,
	}
}
