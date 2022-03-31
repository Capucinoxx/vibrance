package event

import (
	"time"

	"github.com/Capucinoxx/vibrance/server/model"
)

type Repository interface {
	Events(from, to time.Time) (model.Events, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (m repository) Events(from, to time.Time) (model.Events, error) {
	return model.Events{}, nil
}
