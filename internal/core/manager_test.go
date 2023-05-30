package core

import (
	"log"
	"testing"
)

type mockEncryptor struct {
}

// Decrypt implements Encryptor.
func (*mockEncryptor) Decrypt(hash string, data []byte) ([]byte, error) {
	return append([]byte(hash), data...), nil
}

// Encrypt implements Encryptor.
func (*mockEncryptor) Encrypt(hash string, data []byte) ([]byte, error) {
	return append([]byte(hash), data...), nil
}

func Test_dataManager_Set(t *testing.T) {
	tests := []struct {
		name string
		want Setter
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			psw := PasswordForm{
				Password: "password",
				Login:    "login",
				Resource: "https://www.google.com",
			}

			d := NewData()
			err := d.AddEncription(&mockEncryptor{}).
				Set().MetaData("password by google").
				Set().Password("hash", &psw)
			res, _ := d.Get().Data("9090090909090")
			log.Println(res, err)
			// if got := d.Set(); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("dataManager.Set() = %v, want %v", got, tt.want)
			// }
		})
	}
}
