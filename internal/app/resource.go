package app

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/storage/localstore"
)

type Resource interface {
	Login(form *core.LoginForm) error
	Registration(form *core.LoginForm) error
	RegistrationAccept(form *core.RegistrationAccept) error
}

type resource struct {
	store localstore.Store
}

// Login implements Resource.
func (*resource) Login(form *core.LoginForm) error {
	panic("unimplemented")
}

// Registration implements Resource.
func (*resource) Registration(form *core.LoginForm) error {
	panic("unimplemented")
}

// RegistrationAccept implements Resource.
func (*resource) RegistrationAccept(form *core.RegistrationAccept) error {
	panic("unimplemented")
}

func NewResource() Resource {
	return &resource{}
}
