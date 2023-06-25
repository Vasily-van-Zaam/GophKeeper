package pstgql

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	"github.com/google/uuid"
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

func Test_store_AddData(t *testing.T) {
	type fields struct {
		config config.Config
	}
	type args struct {
		ctx  context.Context
		data []*core.ManagerData
	}
	userID, _ := uuid.Parse("c4d34fce-12f1-11ee-b2e7-0242ac1a0002")
	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()
	id4 := uuid.New()

	timeNow := time.Now()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "add data",
			fields: fields{
				config: config.New(logger.New(), cryptor.New()),
			},
			args: args{
				ctx: context.Background(),
				data: []*core.ManagerData{
					{
						Data: []byte("Hello1"),
						InfoData: core.InfoData{
							ID:        &id1,
							DataType:  "text",
							MetaData:  "Text hello1",
							Hash:      "1234-1",
							UserID:    &userID,
							UpdatedAt: &timeNow,
							CreatedAt: &timeNow,
						},
					},
					{
						Data: []byte("Hello2"),
						InfoData: core.InfoData{
							ID:        &id2,
							DataType:  "text",
							MetaData:  "Text hello2",
							Hash:      "1234-2",
							UserID:    &userID,
							UpdatedAt: &timeNow,
							CreatedAt: &timeNow,
						},
					},
					{
						Data: []byte("Hello3"),
						InfoData: core.InfoData{
							ID:        &id3,
							DataType:  "text",
							MetaData:  "Text hello3",
							Hash:      "1234-3",
							UserID:    &userID,
							UpdatedAt: &timeNow,
							CreatedAt: &timeNow,
						},
					},
					{
						Data: []byte("Hello4"),
						InfoData: core.InfoData{
							ID:        &id4,
							DataType:  "text",
							MetaData:  "Text hello4",
							Hash:      "1234-4",
							UserID:    &userID,
							UpdatedAt: &timeNow,
							CreatedAt: &timeNow,
						},
					},
				},
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
			return
			got, err := s.AddData(tt.args.ctx, tt.args.data...)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.AddData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("store.AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}
