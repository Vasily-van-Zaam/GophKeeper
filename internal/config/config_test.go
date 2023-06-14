// Configuration package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/logger"
)

func TestNew(t *testing.T) {
	crypt := cryptor.New()
	type args struct {
		logger Logger
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				logger: logger.New(),
			},
			want: "datastore, 43200, 10, :3200, secret_key_version_0.0.0, secret_key_version_0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := New(tt.args.logger, crypt)

			path := conf.Client().FilePath()
			refresh := conf.Server().Expires(true)
			access := conf.Server().Expires()
			addr := conf.Server().RunAddrss()
			v000 := conf.Server().SecretKey("0.0.0")
			v001 := conf.Server().SecretKey("0.0.1")
			// v002 := conf.Server().SecretKey("0.0.2")
			got := fmt.Sprintf("%v, %v, %v, %v, %v, %v", path, refresh, access, addr, v000, v001)
			// log.Println(v002)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
