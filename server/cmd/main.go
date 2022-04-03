package main

import (
	"log"
	"time"

	"github.com/Capucinoxx/vibrance/internal/pkg/common/router"
	"github.com/Capucinoxx/vibrance/internal/pkg/connector/cassandra"
	"github.com/Capucinoxx/vibrance/internal/pkg/service/client"
	"github.com/Capucinoxx/vibrance/internal/pkg/service/token"
	"github.com/Capucinoxx/vibrance/internal/server"
)

func main() {
	cass, err := cassandra.New(cassandra.Config{
		Hosts:        []string{"127.0.0.1"},
		Port:         9042,
		ProtoVersion: 4,
		Consistency:  "Quorum",
		Keyspace:     "cdb",
		Timeout:      time.Second * 5,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer cass.Close()

	c := router.Consumer("", nil, nil)

	c.AddSubRouter(client.Handle(cass, time.Second))
	c.AddSubRouter(token.Handle(cass, time.Second))

	c.Consume()
	server.Start()
}
