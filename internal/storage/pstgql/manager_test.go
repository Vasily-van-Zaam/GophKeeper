package pstgql

import (
	"context"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func Test_store_GetDataInfo(t *testing.T) {
	type fields struct {
		config config.Config
	}
	type args struct {
		ctx    context.Context
		userID string
		types  []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*core.ManagerData
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				config: config.New(logger.New(), cryptor.New()),
			},
			args: args{
				ctx:    context.Background(),
				userID: "ecfb301e-12c1-11ee-992c-0242ac1a0002",
				types:  []string{"text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New(tt.fields.config)
			if err != nil {
				t.Errorf("store.GetDataInfo() err  = %v", err)
				return
			}
			got, err := s.GetData(tt.args.ctx, tt.args.userID, tt.args.types...)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.GetDataInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.GetDataInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
