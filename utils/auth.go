package utils

type Auth struct {
	token string;
}

// TODO: implement this Generate token method to have a timestamp
func (a *Auth) GenerateToken() {
	a.token = "token"
}

// TODO: implement this method to check if the token is valid
func (a Auth) IsValidToken() bool {
	return false
}

func (a Auth) GetToken() string {
	return a.token
}

func (a *Auth) InvalidToken() {
	a.token = ""
}