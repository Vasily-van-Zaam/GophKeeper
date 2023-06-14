package core

// Model User.
type User struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Age            string `json:"age"`
	Sex            string `json:"sex"`
	Email          string `json:"email"`
	PrivateKey     string `json:"private_key"`
	ID             string `json:"id"`
	EmailConfirmed bool   `json:"email_confirmed"`
}

type UserProfile struct {
	FirrsName string `json:"firrs_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
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
