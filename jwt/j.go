package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	ErrExpired 				= errors.New("jwt: token expired")
	ErrNotValidYet 			= errors.New("jwt: token is not valid yet")
	ErrMalformed 			= errors.New("jwt: token malformed")
	ErrNotCustomClaims 		= errors.New("jwt: token is not custom claims")
	ErrUnknownError 		= errors.New("jwt: unknown error")
)

const AccessTokenSecretKey 	= "ACCESS_TOKEN_SECRET_KEY"
const RefreshTokenSecretKey = "REFRESH_TOKEN_SECRET_KEY"

type JWT interface {
	Verify(token string) error
	VerifyCustomClaims(token string) (*CustomClaims, error)
	SignIn(exp time.Duration, username string, email *string, roles []string, userId *int64) (string, error)
	Sign(method jwt.SigningMethod, claims jwt.Claims)   (string, error)
}

type microkitjwt struct {
	secret []byte
}

func (kit microkitjwt) Verify(t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return kit.secret, nil
	})
	if err := kit.isError(err); err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	log.Print("jwt: is not token err: "+err.Error())
	return ErrUnknownError
}

func (kit microkitjwt) isError(err error) error {
	if err == nil {
		return nil
	}
	ve, isValidationError := err.(*jwt.ValidationError)
	if !isValidationError {
		log.Print("jwt: unknown error: "+err.Error())
		return ErrUnknownError
	}
	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		return ErrMalformed
	} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
		return ErrNotValidYet
	} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
		return ErrExpired
	}
	log.Print("jwt: is not token err: "+err.Error())
	return ErrUnknownError

}

func (kit microkitjwt) VerifyCustomClaims(t string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(t, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return kit.secret, nil
	})
	if err := kit.isError(err); err != nil {
		return nil, err
	}
	if token.Valid {
		claims, isCustomClaims := token.Claims.(*CustomClaims)
		if !isCustomClaims {
			return nil, ErrNotCustomClaims
		}
		return claims, nil
	}
	log.Print("jwt: for unknown reason token is not valid")
	return nil, ErrUnknownError
}

func (kit microkitjwt) Sign(method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(kit.secret)
}

func (kit microkitjwt) SignIn(exp time.Duration, username string, email *string, roles []string, userId *int64) (string, error) {
	claims := CustomClaims{
		Username:       username,
		Email:          email,
		Roles:          roles,
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
	}
	return kit.Sign(jwt.SigningMethodHS256, claims)
}

func NewJWT(secret []byte) JWT {
	return microkitjwt{secret:secret}
}

type CustomClaims struct {
	Username string 	`json:"username"`
	Email 	 *string 	`json:"email"`
	Roles 	 []string  	`json:"roles"`
	UserId   *int64     `json:"user_id"`
	jwt.StandardClaims
}