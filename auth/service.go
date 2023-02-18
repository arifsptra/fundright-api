package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
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

// function to validate token with jwt
func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	// check token validity
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		// checks if the method used to generate the token signature is the HMAC method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtSecretKey), nil
	})

	// if an error occurs in the token validation process
	if err != nil {
		return token, err
	}

	// if token not valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// if token valid
	return token, nil
}
