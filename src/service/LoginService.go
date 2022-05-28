package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type LoginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		email:    "login@email.com",
		password: "abcd1234",
	}
}

func (info *LoginInformation) LoginUser(email string password string)bool{
	return info.email == email && info.password == password
}
