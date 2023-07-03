package dddcore

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

// JwtTokenParser represents a parser for JWT tokens.
type JwtTokenParser *jwt.Parser

// JwtToken represents a parsed JWT token.
type JwtToken *jwt.Token

// NewJwtTokenParser creates a new instance of JwtTokenParser with the specified token validation methods.
// It uses the jwt.NewParser function from the "github.com/golang-jwt/jwt/v5" package to create the parser.
func NewJwtTokenParser(methods []string) JwtTokenParser {
	return jwt.NewParser(
		jwt.WithValidMethods(methods),
	)
}

// JwtParse parses a JWT token using the provided parser, token string, and key.
// It returns a JwtToken if the token is valid and the claims are successfully parsed.
// It performs additional checks on the token claims, such as verifying the expiration time (exp) and issued at time (iat).
// If any validation checks fail or the claims are invalid, appropriate error messages are returned.
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
