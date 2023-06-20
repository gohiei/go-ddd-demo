package dddcore

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtTokenParser *jwt.Parser
type JwtToken *jwt.Token

func NewJwtTokenParser(methods []string) JwtTokenParser {
	return jwt.NewParser(
		jwt.WithValidMethods(methods),
	)
}

func JwtParse(parser *jwt.Parser, token string, key interface{}) (JwtToken, error) {
	tt, err := parser.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tt.Claims.(*jwt.RegisteredClaims); ok && tt.Valid {
		if exp, err := claims.GetExpirationTime(); exp == nil || err != nil {
			return nil, errors.New("invalid exp")
		}

		if iat, err := claims.GetIssuedAt(); iat == nil || err != nil {
			return nil, errors.New("invalid iat")
		}

		return tt, nil
	}

	return nil, errors.New("invalid claims")
}
