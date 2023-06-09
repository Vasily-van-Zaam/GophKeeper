package core

// Model User.
type User struct {
	FirrsName string `json:"firrs_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	ID        string `json:"id"`
}

type UserProfile struct {
	FirrsName      string `json:"firrs_name"`
	LastName       string `json:"last_name"`
	Age            string `json:"age"`
	Sex            string `json:"sex"`
	Email          string `json:"email"`
	EmailConfirmed bool   `json:"email_confirmed"`
}

// User Login form for authenticated.
type LoginForm struct {
	Email   string `json:"email"`
	Pasword string `json:"pasword"`
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
