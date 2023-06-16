// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
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
