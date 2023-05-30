package service

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
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
