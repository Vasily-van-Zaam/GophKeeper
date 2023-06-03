package localstore

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	logg := logger.New()
	tests := []struct {
		name    string
		config  config.Config
		want    Store
		wantErr bool
	}{
		{
			name: "new store",
			want: &store{
				data: &core.DataGob{},
			},
			config: config.New(logg),
		},
		{
			name:    "new store error",
			want:    &store{},
			wantErr: true,
			config:  config.New(logg),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Remove(tt.config.Client().FilePath())
			if err != nil {
				log.Println(err)
			}
			got, err := New(tt.config)
			if err != nil && !tt.wantErr {
				t.Errorf("New() = err %v, want %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_store_GetData(t *testing.T) {
	type fields struct {
		data *core.DataGob
	}
	userID1 := uuid.New()
	userID2 := uuid.New()
	userID3 := uuid.New()
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
			name: "store_get_data found user1",
			fields: fields{
				data: &core.DataGob{
					DataList: []*core.ManagerData{
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID1,
							},
						},
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID2,
							},
						},
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: userID1.String(),
			},
			want: []*core.ManagerData{
				{
					Data: nil,
					InfoData: core.InfoData{
						UserID: &userID1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "store_get_data not found",
			fields: fields{
				data: &core.DataGob{
					DataList: []*core.ManagerData{
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID1,
							},
						},
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID2,
							},
						},
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: userID3.String(),
			},
			want:    []*core.ManagerData{},
			wantErr: false,
		},
		{
			name: "store_get_data found user2",
			fields: fields{
				data: &core.DataGob{
					DataList: []*core.ManagerData{
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID1,
							},
						},
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &userID2,
							},
						},
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: userID2.String(),
			},
			want: []*core.ManagerData{
				{
					Data: nil,
					InfoData: core.InfoData{
						UserID: &userID2,
					},
				},
			},
			wantErr: false,
		},
		{
			name:   "store_get_data_error",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				userID: userID2.String(),
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				data: tt.fields.data,
			}
			got, err := s.GetData(tt.args.ctx, tt.args.userID, tt.args.types...)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_AddData(t *testing.T) {
	type fields struct {
		data *core.DataGob
	}
	type args struct {
		ctx  context.Context
		data *core.ManagerData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.ManagerData
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				data: tt.fields.data,
			}
			got, err := s.AddData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.AddData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_saveToFile(t *testing.T) {
	type fields struct {
		data     *core.DataGob
		filePath string
		config   config.Config
	}
	id := uuid.New()
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "store_save_to_file",
			fields: fields{
				data: &core.DataGob{
					DataList: []*core.ManagerData{
						{
							Data: nil,
							InfoData: core.InfoData{
								UserID: &id,
							},
						},
					},
				},
				config:   config.New(logger.New()),
				filePath: "test_store_save_to_file",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				data:     tt.fields.data,
				filePath: tt.fields.filePath,
				config:   tt.fields.config,
			}
			if err := s.saveToFile(); (err != nil) != tt.wantErr {
				t.Errorf("store.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
