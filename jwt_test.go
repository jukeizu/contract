package contract

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

func TestGenerateSignedTokenString(t *testing.T) {
	claims := JwtClaims{
		ChannelId: "channelId",
		AuthorId:  "authorId",
		ServerId:  "serverId",
		Platform:  "platform",
	}

	signedString, err := claims.SignedTokenString(secret)
	if err != nil {
		t.Error(err)
	}

	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		t.Error(err)
	}

	if !token.Valid {
		t.Errorf("did not generate a valid token")
	}

	t.Log(signedString)
}

func TestParseClaims(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhaWQiOiJhdXRob3JJZCIsImNpZCI6ImNoYW5uZWxJZCIsInNpZCI6InNlcnZlcklkIiwicGx0IjoicGxhdGZvcm0ifQ.P_yCxY1txkcgn4wD8UsMIYQoMhCKWadAnpCNmfjtaOU"
	claims, err := ParseJwtClaims(token, secret)
	if err != nil {
		t.Error(err)
	}

	want := "platform"
	have := claims.Platform

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}
