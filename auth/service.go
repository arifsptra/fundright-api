package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{}

// initialize the secret key for the jwt service
var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func NewService() *jwtService {
	return &jwtService{}
}

// function to generate token with jwt
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// header jwt token
	header := jwt.SigningMethodHS256
	
	// payload jwt token
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	
	// create jwt with claim
	token := jwt.NewWithClaims(header, payload)

	// sign JWT with secret key
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}