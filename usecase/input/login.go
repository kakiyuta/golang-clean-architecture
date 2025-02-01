package input

type Login struct {
	Email    string
	Password string
}

func NewLoginRequestUser(email, password string) *Login {
	return &Login{
		Email:    email,
		Password: password,
	}
}
