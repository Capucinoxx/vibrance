package client

import (
	"context"
	"time"

	"github.com/Capucinoxx/vibrance/internal/pkg/common/oauth"
	"github.com/Capucinoxx/vibrance/internal/pkg/connector/cassandra"
	"github.com/gocql/gocql"
)

type Repository interface {
	Create(ctx context.Context, client oauth.Client) error
	Find(ctx context.Context, key string) (oauth.Client, error)
	UpdateSecret(ctx context.Context, key, secret string) error
	SoftDelete(ctx context.Context, key string) error
	Delete(ctx context.Context, key string) error
}

type repository struct {
	conn    *gocql.Session
	timeout time.Duration
}

func NewRepository(conn *gocql.Session, timeout time.Duration) Repository {
	return &repository{conn, timeout}
}

func (r repository) Create(ctx context.Context, client oauth.Client) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	applied, err := r.conn.Query(
		StmtCreateClient,
		client.ID,
		client.Key,
		client.Secret,
		client.CreatedAt,
		client.DeletedAt,
	).WithContext(ctx).MapScanCAS(map[string]interface{}{})

	if err != nil {
		return err
	}

	if !applied {
		return cassandra.ErrDuplication
	}

	return nil
}

func (r repository) Find(ctx context.Context, key string) (oauth.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var client oauth.Client

	if err := r.conn.Query(StmtFindClient, key).WithContext(ctx).Scan(
		&client.ID,
		&client.Key,
		&client.Secret,
		&client.CreatedAt,
		&client.DeletedAt,
	); err != nil {
		if err == gocql.ErrNotFound {
			err = cassandra.ErrNotFound
		}
		return oauth.Client{}, err
	}

	return client, nil
}

func (r repository) UpdateSecret(ctx context.Context, key, secret string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	applied, err := r.conn.Query(
		StmtUpdateSecretClient,
		secret,
		key,
	).WithContext(ctx).MapScanCAS(map[string]interface{}{})

	if err != nil {
		return err
	}

	if !applied {
		return cassandra.ErrNotFound
	}

	return nil
}

func (r repository) SoftDelete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	applied, err := r.conn.Query(
		StmtSoftDeleteClient,
		time.Now().UTC(),
		key,
	).WithContext(ctx).MapScanCAS(map[string]interface{}{})

	if err != nil {
		return err
	}

	if !applied {
		return cassandra.ErrNotFound
	}

	return nil
}

func (r repository) Delete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	applied, err := r.conn.Query(
		StmtDeleteClient,
		key,
	).WithContext(ctx).MapScanCAS(map[string]interface{}{})

	if err != nil {
		return err
	}

	if !applied {
		return cassandra.ErrNotFound
	}

	return nil
}
