package oauth

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/gocql/gocql"
)

func GenerateToken(key string, ttl TokenTTL, scopes []TokenScope) Token {
	secret := gocql.TimeUUID().String()
	h := sha1.New()
	h.Write([]byte(key))

	return Token{
		ClientKey:    key,
		ClientSecret: secret,
		Hash:         hex.EncodeToString(h.Sum([]byte(secret))),
		TTL:          ttl,
		Scopes:       scopes,
	}
}
