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
	FirrsName string `json:"firrs_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
}

// User Login form for authenticated.
type LoginForm struct {
	Email   string `json:"email"`
	Pasword string `json:"pasword"`
}

// AuthToken access + refrech token for authenticated.
type AuthToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
