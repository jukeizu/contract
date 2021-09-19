package contract

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Request struct {
	Id          string              `json:"id,omitempty" yaml:"id,omitempty"`
	ContextId   string              `json:"contextId,omitempty" yaml:"contextId,omitempty"`
	IntentId    string              `json:"intentId,omitempty" yaml:"intentId,omitempty"`
	Source      string              `json:"source,omitempty" yaml:"source,omitempty"`
	Bot         User                `json:"bot,omitempty" yaml:"bot,omitempty"`
	Author      User                `json:"author,omitempty" yaml:"author,omitempty"`
	ChannelId   string              `json:"channelId,omitempty" yaml:"channelId,omitempty"`
	ServerId    string              `json:"serverId,omitempty" yaml:"serverId,omitempty"`
	Servers     []Server            `json:"servers,omitempty" yaml:"servers,omitempty"`
	Mentions    []User              `json:"mentions,omitempty" yaml:"mentions,omitempty"`
	Content     string              `json:"content,omitempty" yaml:"content,omitempty"`
	Application Application         `json:"application,omitempty" yaml:"application,omitempty"`
	QueryParams map[string][]string `json:"queryParams,omitempty"`
}

type User struct {
	Id            string `json:"id,omitempty" yaml:"id,omitempty"`
	Name          string `json:"name,omitempty" yaml:"name,omitempty"`
	Discriminator string `json:"discriminator" yaml:"discriminator"`
}

type Server struct {
	Id              string `json:"id,omitempty" yaml:"id,omitempty"`
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerId         string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	UserCount       int32  `json:"userCount,omitempty" yaml:"userCount,omitempty"`
	IconUrl         string `json:"iconUrl,omitempty" yaml:"iconUrl,omitempty"`
	SystemChannelId string `json:"systemChannelId,omitempty" yaml:"systemChannelId,omitempty"`
}

type Application struct {
	Id          string `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Icon        string `json:"icon,omitempty" yaml:"icon,omitempty"`
	Owner       User   `json:"owner,omitempty" yaml:"owner,omitempty"`
}

// Server returns the full server for the ServerId
func (r Request) Server() Server {
	for _, server := range r.Servers {
		if server.Id == r.ServerId {
			return server
		}
	}

	return Server{Id: r.ServerId}
}

// JwtClaims returns jwt claims populated with values from request
func (r Request) JwtClaims(expiresIn time.Duration) JwtClaims {
	return JwtClaims{
		ContextId: r.ContextId,
		IntentId:  r.IntentId,
		ChannelId: r.ChannelId,
		AuthorId:  r.Author.Id,
		ServerId:  r.ServerId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
}
