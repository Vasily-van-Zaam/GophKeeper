package cryptor

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_generateRandom(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "genarate random ok",
			args: args{
				size: 12,
			},
			want: 12,
		},
		{
			name: "genarate random ok",
			args: args{
				size: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandom(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("generateRandom() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_cryptor_Encrypt(t *testing.T) {
	type args struct {
		secret        []byte
		secretDecrypt []byte
		data          []byte
	}
	tests := []struct {
		name    string
		c       cryptor
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "cryptor_encrypt",
			args: args{
				secret:        []byte("master_password"),
				data:          []byte("test text"),
				secretDecrypt: []byte("master_password"),
			},
			want:    []byte("test text"),
			wantErr: false,
		},
		{
			name: "cryptor_encrypt",
			args: args{
				secret:        []byte("master_password"),
				data:          []byte("test text"),
				secretDecrypt: []byte("master_password_error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got, _ := c.Encrypt(tt.args.secret, tt.args.data)

			got2, err2 := c.Decrypt(tt.args.secretDecrypt, got)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("cryptor.Decrypt() error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got2) {
				t.Errorf("cryptor.Decrypt() = %v, want %v", got2, tt.want)
			}
		})
	}
}

func Test_cryptor_Decrypt(t *testing.T) {
	type args struct {
		secret []byte
		data   []byte
	}
	data, err := hex.DecodeString("ca634228d0d53f951ec8648b7692d22be98263cf7f511c21668d34e2c6220317c19a8fb3d4")
	if err != nil {
		t.Errorf("cryptor.Decrypt() error %v", err)
		return
	}
	tests := []struct {
		name    string
		c       cryptor
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "decrypt",
			args: args{
				secret: []byte("master_password"),
				data:   data,
			},
			want: []byte{116, 101, 115, 116, 32, 116, 101, 120, 116},
		},
		{
			name: "decrypt",
			args: args{
				secret: []byte("master_password1"),
				data:   data,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got, err := c.Decrypt(tt.args.secret, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("cryptor.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cryptor.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
