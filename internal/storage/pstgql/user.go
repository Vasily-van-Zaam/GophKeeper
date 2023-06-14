package pstgql

import (
	"context"
	"errors"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

// GetUserByEmail implements Store.
func (*store) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	if email == "test@mail.ru" {
		return &core.User{
			Email:      "test@mail.ru",
			PrivateKey: "private_user_key",
		}, nil
	}
	return nil, errors.New("not found")
}
