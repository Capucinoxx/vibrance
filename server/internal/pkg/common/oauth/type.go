package oauth

import "time"

type TokenScope string

const (
	TokenAllGrant  TokenScope = ""
	TokenReadGrant TokenScope = "read"
)

type TokenTTL int

const (
	TokenTTLAccess  TokenTTL = 3600
	TokenTTLRefresh TokenTTL = 86400
)

type Client struct {
	ID        string
	Key       string
	Secret    string
	CreatedAt time.Time
	DeletedAt *time.Time
}

type Token struct {
	ClientKey    string
	ClientSecret string
	Hash         string
	TTL          TokenTTL
	Scopes       []TokenScope
}
