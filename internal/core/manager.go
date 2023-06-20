package core

import (
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func setVersionHash(b []byte) string {
	hash := sha512.New()
	hash.Write(b)
	res := fmt.Sprintf("%x", hash.Sum(nil))
	return res
}

// The data manager model for saving.
type manager struct {
	infoData  *InfoData
	data      []byte // encrypted
	setter    Setter
	getter    Getter
	encryptor Encryptor
}

// ToData implements Manager.
func (d *manager) ToData() (*ManagerData, error) {
	if d.data == nil {
		return nil, errors.New("data is nil")
	}
	if d.infoData == nil {
		return nil, errors.New("infoData is nil")
	}
	return &ManagerData{
		Data:     d.data,
		InfoData: *d.infoData,
	}, nil
}

type setter struct {
	data *manager
}

// TryPassword implements Setter.
func (s *setter) TryPassword(masterPsw string, count int) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(count)
	if err != nil {
		return err
	}
	return s.setData(masterPsw, DataTypeTryPassword, d)
}

func (s *setter) isCreating() bool {
	if s.data != nil {
		if s.data.infoData != nil {
			if s.data.infoData.CreatedAt != nil {
				return false
			}
		}
	}
	return true
}
func (s *setter) setData(masterPsw string, dType DataType, d []byte) error {
	bs, err := s.data.encryptor.Encrypt([]byte(masterPsw), d)
	if err != nil {
		return err
	}
	nowTime := time.Now().UTC()

	s.data.infoData.UpdatedAt = &nowTime
	if s.isCreating() {
		s.data.infoData.CreatedAt = &nowTime
	}
	s.data.infoData.Hash = setVersionHash(d)
	s.data.infoData.DataType = string(dType)
	s.data.data = bs
	return nil
}

// BankCard implements Setter.
func (s *setter) BankCard(masterPsw string, card *BankCardFomm) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(card)
	if err != nil {
		return err
	}
	return s.setData(masterPsw, DataTypeCard, d)
}

// Password implements Setter.
func (s *setter) Password(masterPsw string, psw *PasswordForm) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(psw)
	if err != nil {
		return err
	}
	return s.setData(masterPsw, DataTypePassword, d)
}

// Text implements Setter.
func (s *setter) Text(masterPsw string, text string) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d := []byte(text)
	return s.setData(masterPsw, DataTypeText, d)
}

// Data implements Setter.
func (s *setter) Data(masterPsw string, data []byte) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	return s.setData(masterPsw, DataTypeFile, data)
}

// Data implements Setter.
func (s *setter) AccessData(masterPsw string, user *User) error {
	if s.data.encryptor == nil {
		return errors.New("need add encrypter")
	}
	d, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.setData(masterPsw, DataTypeUser, d)
}

// MetaData implements Setter.
func (s *setter) MetaData(metaData string) Manager {
	s.data.infoData.MetaData = metaData
	return s.data
}

type getter struct {
	data *manager
}

// EncryptData implements Getter.
func (g *getter) EncryptData() []byte {
	return g.data.data
}

// Data implements Getter.
func (g *getter) Data(masterPsw string) ([]byte, error) {
	if g.data.encryptor == nil {
		return nil, errors.New("need add encrypter")
	}
	if g.data.data == nil {
		return nil, errors.New("data is nil")
	}
	bs, err := g.data.encryptor.Decrypt([]byte(masterPsw), g.data.data)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// InfoData implements Getter.
func (g *getter) InfoData() *InfoData {
	return g.data.infoData
}

// AddEncription implements Manager.
func (d *manager) AddEncription(encrypt Encryptor) Manager {
	d.encryptor = encrypt
	return d
}

// Get implements Manager.
func (d *manager) Get() Getter {
	return d.getter
}

// Set implements Manager.
func (d *manager) Set() Setter {
	return d.setter
}

// Create new manager for store.
func NewManager() Manager {
	data := &manager{
		infoData: &InfoData{},
	}
	data.getter = &getter{
		data: data,
	}
	data.setter = &setter{
		data: data,
	}

	return data
}

// Create manager form storage.
func NewManagerFromData(data *ManagerData) Manager {
	d := &manager{
		infoData: &data.InfoData,
		data:     data.Data,
	}
	d.getter = &getter{
		data: d,
	}
	d.setter = &setter{
		data: d,
	}
	return d
}
