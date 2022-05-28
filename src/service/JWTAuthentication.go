package service

type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name,omitempty"`
	User bool   `json:"user,omitempty"`
	jwt.StandardClaims
}


type JWTServices struct {
	secretKey string
	issuer string

  
}

func JWSAuthService()JWTService {
	return & *jwtServices{
		secretKey: getSecretKey()
		issuer : "charles",
	}
}

func getSecretKey(){
	secret := os.Getenv("SECRET")
	if secret== "" {
		secret = "secret"

	}
	return secret
}

func (Service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		emails,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix()
			Issuer : service.issuer,
			issuedAt: time.Now().Unix()
		},		
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err :=token.SignedString([]byte(Service.secretKey))
	if err != nil {
		panic(err)
	}

	return(t)
}