package core

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/Vasily-van-Zaam/GophKeeper.git/pkg/cryptor"
	"github.com/google/uuid"
)

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
			cryp := cryptor.New()
			uID := uuid.New()
			d := NewManager(&uID)
			err1 := d.AddEncription(cryp).
				Set().MetaData("password by google").
				Set().Password("password", &psw)
			res, err := d.Get().Data("password")

			log.Println(string(res), err, err1)
			log.Println(hex.EncodeToString(d.Get().EncryptData()))
			// saveData := "7a82c2a4635b18f44d811d8d64d4eb180af5cb8b20d4c23e4c59006b80fbb41c318ce3c4aed7b5192f0cd23e1216bd8dc17320c00047f15f70114381abe53bc2e4fe542d7b422e9067f10f946acfcc6115c82bb0df7bcbf99a5259ba9ce17da35de784c3"

			// strDD, _ := hex.DecodeString(saveData)
			dd := ManagerData{
				Data:     d.Get().EncryptData(), // []byte(strDD),
				InfoData: *d.Get().InfoData(),
			}
			m := dd.ToManager()

			m.AddEncription(cryp)

			// bdd, _ := json.Marshal(dd)
			getData, e := m.Get().Data("password")
			log.Println(string(getData), e)

			log.Println(m.ToData())
			// log.Println(string(res), err, err1)
			// log.Println(hex.EncodeToString(d.Get().EncryptData()))
			// if got := d.Set(); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("dataManager.Set() = %v, want %v", got, tt.want)
			// }
		})
	}
}
