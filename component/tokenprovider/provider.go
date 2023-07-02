package tokenprovider

import (
	"MyPetProject/commons"
	"errors"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = commons.NewCustomError(
		errors.New("Token not found"), "token not found", "ErrNotFound")
	ErrEncodingToken = commons.NewCustomError(errors.New("error encoding the token"), "error encoding the token", "ErrEncodingToken")
	ErrInvalidToken  = commons.NewCustomError(errors.New("invalid token provided"), "invalid token provided", "ErrInvalidToken")
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

// /
type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
