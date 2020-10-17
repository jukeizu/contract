package contract

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims represents custom jwt claims for treediagram
type JwtClaims struct {
	jwt.StandardClaims
	ContextId string `json:"ctx,omitempty"`
	IntentId  string `json:"iid,omitempty"`
	AuthorId  string `json:"aid,omitempty"`
	ChannelId string `json:"cid,omitempty"`
	ServerId  string `json:"sid,omitempty"`
	Platform  string `json:"plt,omitempty"`
}

// Token provides access to the underlying jwt.Token
func (c JwtClaims) Token() *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c)
}

// SignedTokenString generates a signed jwt token string
func (c JwtClaims) SignedTokenString(secret []byte) (string, error) {
	return c.Token().SignedString(secret)
}

// ParseJwtClaims parses claims from and validates a string token
func ParseJwtClaims(tokenString string, secret []byte) (JwtClaims, error) {
	claims := JwtClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return claims, err
}
