package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const jwtSecretKey = "1234567890123456789012"

// Payload contains the payload data of the struct
type Payload struct {
	ID        uuid.UUID
	Username  string    `json:"username"`
	IsLogin   bool      `json:"is_login"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// GenerateToken generates a new token for a specific username and duration
func GenerateToken(username string, duration time.Duration) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IsLogin:   true,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

// Valid checks if the token is expired or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// VerifyToken checks if the token is valid or not
func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
