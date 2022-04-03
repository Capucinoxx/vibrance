package token

import (
	"context"
	"time"

	"github.com/Capucinoxx/vibrance/internal/pkg/common/oauth"
	"github.com/Capucinoxx/vibrance/internal/pkg/connector/cassandra"
	"github.com/gocql/gocql"
)

type Repository interface {
	Create(ctx context.Context, accTok oauth.Token, refTok oauth.Token) error
	Refresh(ctx context.Context, refTokHash string, accTok oauth.Token, refTok oauth.Token) error
	Revoke(ctx context.Context, hash string) error
	Find(ctx context.Context, hash string) (oauth.Token, error)
}

type repository struct {
	conn    *gocql.Session
	timeout time.Duration
}

func NewRepository(conn *gocql.Session, timeout time.Duration) Repository {
	return &repository{conn, timeout}
}

func (r repository) Create(ctx context.Context, accTok, refTok oauth.Token) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	batch := r.conn.NewBatch(gocql.LoggedBatch).WithContext(ctx)
	batch.Query(StmtCreateToken, accTok.Hash, accTok.ClientKey, accTok.ClientSecret, accTok.Scopes, accTok.TTL)
	batch.Query(StmtCreateToken, refTok.Hash, refTok.ClientKey, refTok.ClientSecret, refTok.Scopes, refTok.TTL)

	return r.conn.ExecuteBatch(batch)
}
func (r repository) Refresh(ctx context.Context, refTokHash string, accTok, refTok oauth.Token) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	batch := r.conn.NewBatch(gocql.LoggedBatch).WithContext(ctx)
	batch.Query(StmtDeleteToken, refTokHash)
	batch.Query(StmtCreateToken, accTok.Hash, accTok.ClientKey, accTok.ClientSecret, accTok.Scopes, accTok.TTL)
	batch.Query(StmtCreateToken, refTok.Hash, refTok.ClientKey, refTok.ClientSecret, refTok.Scopes, refTok.TTL)

	return r.conn.ExecuteBatch(batch)
}
func (r repository) Revoke(ctx context.Context, hash string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	return r.conn.Query(StmtDeleteToken, hash).WithContext(ctx).Exec()
}

func (r repository) Find(ctx context.Context, hash string) (oauth.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var token oauth.Token

	err := r.conn.Query(StmtFindToken, hash).WithContext(ctx).Scan(
		&token.Hash,
		&token.ClientKey,
		&token.ClientSecret,
		&token.Scopes,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			err = cassandra.ErrNotFound
		}
		return oauth.Token{}, err
	}

	return token, nil
}
