package cassandra

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
)

var (
	ErrDuplication = errors.New("duplicated record")
	ErrNotFound    = errors.New("record not found")
)

type Config struct {
	Hosts        []string
	Port         int
	ProtoVersion int
	Consistency  string
	Keyspace     string
	Timeout      time.Duration
}

func New(config Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Hosts...)

	cluster.Port = config.Port
	cluster.ProtoVersion = config.ProtoVersion
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.ParseConsistency(config.Consistency)
	cluster.Timeout = config.Timeout

	return cluster.CreateSession()
}
