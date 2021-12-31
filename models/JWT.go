package models

import (
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

const (
	SessionIdKey = "sid"
)

type JWT struct {
	SessionId string
	signKey   []byte
	algorithm jwa.SignatureAlgorithm
}

func NewJWT(signKey []byte, algorithm jwa.SignatureAlgorithm) *JWT {
	return &JWT{
		signKey:   signKey,
		algorithm: algorithm,
	}
}

//Encode creates a new JWT, adds in the claims, and signs the token.
//It returns signed token string.
func (j *JWT) Encode(sessionId string, expiry time.Time) (string, error) {
	t, err := jwt.NewBuilder().
		IssuedAt(time.Now().UTC()).
		Expiration(expiry.UTC()).
		Claim(SessionIdKey, sessionId).
		Build()
	if err != nil {
		return "", err
	}
	payload, err := jwt.Sign(t, j.algorithm, j.signKey)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

//Decode parses and verifies a JWT
func (j *JWT) Decode(tokenString string) (jwt.Token, error) {
	t, err := jwt.Parse([]byte(tokenString), jwt.WithVerify(j.algorithm, j.signKey))
	if err != nil {
		return nil, err
	}

	if t == nil {
		return nil, errors.New("failed to parse token")
	}

	if err := jwt.Validate(t); err != nil {
		return nil, err
	}

	return t, nil
}
