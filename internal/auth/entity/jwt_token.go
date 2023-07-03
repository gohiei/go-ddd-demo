package auth

import (
	"cypt/internal/dddcore"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"

	event "cypt/internal/auth/entity/events"
)

var (
	tokenPrefix = "Bearer "
	tokenKey    = "this.is.a.token"
)

type JwtToken struct {
	dddcore.AggregateRoot
	token   string
	parser  *jwt.Parser
	request Request
}

type Request struct {
	Method string
	URL    string
	IP     string
	XFF    string
}

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

func (t *JwtToken) Valid() bool {
	var valid bool
	var jwtErr error

	checked := false

	len := len(tokenPrefix)
	prefix := t.token[:len]
	token := t.token[len:]

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
		t.AddDomainEvent(event.NewInvalidRequestOccurredEvent(
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
