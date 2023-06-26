package core

import "github.com/google/uuid"

// Model User.
type User struct {
	Email          string     `json:"email"`
	PrivateKey     string     `json:"private_key"`
	ID             *uuid.UUID `json:"id"`
	EmailConfirmed bool       `json:"email_confirmed"`
}

// User Login form for authenticated.
type AccessForm struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Token []byte `json:"token"`
}

type RegistrationAccept struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// AuthToken access + refrech token for authenticated.
type AuthToken struct {
	Access  []byte `json:"access"`
	Refresh []byte `json:"refresh"`
	UserKey []byte `json:"userKey"`
}
