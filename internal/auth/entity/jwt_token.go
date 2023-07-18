// Package entity provides entities related to authentication functionality.
package entity

import (
	"net/http"

	"cypt/internal/auth/entity/events"
	"cypt/internal/dddcore"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Constants
var (
	tokenPrefix = "Bearer "
	tokenKey    = "this.is.a.token"
)

// JwtToken represents a JWT token.
type JwtToken struct {
	dddcore.AggregateRoot
	token   string
	parser  *jwt.Parser
	request Request
}

// Request represents an HTTP request.
type Request struct {
	Method string
	URL    string
	IP     string
	XFF    string
}

// NewJwtToken creates a new JwtToken instance.
func NewJwtToken(token string, req Request) (*JwtToken, error) {
	if len(token) < len(tokenPrefix) {
		return nil, dddcore.NewErrorS("00001", "invalid token", http.StatusForbidden)
	}

	return &JwtToken{
		AggregateRoot: dddcore.NewAggregateRoot(),
		token:         token,
		parser:        dddcore.NewJwtTokenParser([]string{"HS256", "HS384", "HS512"}),
		request:       req,
	}, nil
}

// Valid checks if the JWT token is valid.
func (t *JwtToken) Valid() bool {
	var valid bool
	var jwtErr error

	checked := false

	length := len(tokenPrefix)
	prefix := t.token[:length]
	token := t.token[length:]

	if prefix != tokenPrefix {
		valid = false
		checked = true
	}

	if !checked {
		valid = true
		checked = true

		_, err := dddcore.JwtParse(t.parser, token, []byte(tokenKey))

		if err != nil {
			valid = false
			jwtErr = err
		}
	}

	if !valid {
		t.AddDomainEvent(events.NewInvalidRequestOccurredEvent(
			t.token,
			t.request.Method,
			t.request.URL,
			t.request.IP,
			t.request.XFF,
			jwtErr,
		))
	}

	return valid
}
