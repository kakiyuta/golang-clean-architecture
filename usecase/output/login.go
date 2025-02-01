package output

// AuthLogin is the output struct for Login
type AuthLogin struct {
	Token string
}

// NewAuthLogin is the constructor for AuthLogin
func NewAuthLogin(token string) *AuthLogin {
	return &AuthLogin{
		Token: token,
	}
}
