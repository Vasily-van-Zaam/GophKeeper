// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

func Test_client_checkValidPassword(t *testing.T) {
	type fields struct {
		pages      *tview.Pages
		app        *tview.Application
		repository repository.Repository
		config     config.Config
	}
	type args struct {
		psw string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test_password_err",
			args: args{
				psw: "password",
			},
			want: false,
		},
		{
			name: "test_password_ok",
			args: args{
				psw: "Dpassword1a!",
			},
			want: true,
		},
		{
			name: "test_password_err",
			args: args{
				psw: "password1a!",
			},
			want: false,
		},
		{
			name: "test_password_err_len",
			args: args{
				psw: "pA!aa13",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				pages:      tt.fields.pages,
				app:        tt.fields.app,
				repository: tt.fields.repository,
				config:     tt.fields.config,
			}
			if got := c.checkValidPassword(tt.args.psw); got != tt.want {
				t.Errorf("client.checkValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CompareDataSync(t *testing.T) {
	type args struct {
		local  []*core.ManagerData
		remote []*core.ManagerData
	}

	id1, _ := uuid.Parse("c57b60e9-6e3b-4339-a0be-745b37616652")
	id2, _ := uuid.Parse("c57b60e9-6e3b-4339-a0be-745b37616651")
	id3, _ := uuid.Parse("c57b60e9-6e3b-4339-a0be-745b37616650")
	id4, _ := uuid.Parse("c57b60e9-6e3b-4339-a0be-745b37616649")
	tests := []struct {
		name string

		args  args
		want  []*core.CopmareData
		want1 []*core.ManagerData
		want2 []*core.ManagerData
	}{
		{
			name: "test compare",
			args: args{
				local: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "123456789",
						},
					},
				},
				remote: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "12345678",
						},
					},
				},
			},
			want: []*core.CopmareData{
				{
					Local: &core.ManagerData{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "123456789",
						},
					},
					Remote: &core.ManagerData{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "12345678",
						},
					},
				},
			},
			want1: []*core.ManagerData{},
			want2: []*core.ManagerData{},
		},
		{
			name: "test compare",
			args: args{
				local: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "12345678",
						},
					},
				},
				remote: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "12345678",
						},
					},
				},
			},
			want:  []*core.CopmareData{},
			want1: []*core.ManagerData{},
			want2: []*core.ManagerData{},
		},
		{
			name: "test compare",
			args: args{
				local: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345-",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
				},
				remote: []*core.ManagerData{
					{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id2,
							Hash: "123456",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id3,
							Hash: "1234567",
						},
					},
					{
						InfoData: core.InfoData{
							ID:   &id4,
							Hash: "12345678",
						},
					},
				},
			},
			want: []*core.CopmareData{
				{
					Local: &core.ManagerData{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345-",
						},
					},
					Remote: &core.ManagerData{
						InfoData: core.InfoData{
							ID:   &id1,
							Hash: "12345",
						},
					},
				},
			},
			want1: []*core.ManagerData{
				{
					InfoData: core.InfoData{
						ID:   &id4,
						Hash: "12345678",
					},
				},
			},
			want2: []*core.ManagerData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{}
			got, got1, got2 := c.CompareDataSync(tt.args.local, tt.args.remote)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CompareDataSync() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("client.CompareDataSync() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("client.CompareDataSync() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
