package application

import (
	"log"
	"ml-x-men/internal/domain"
)

type Handler interface {
	PersonLogic
	StatLogic
}

type PersonLogic interface {
	GetByDna(dna string) (*domain.Person, error)
	IsMutant(matrix [][]byte) (bool, error)
	PersonCreate(dna string, isMutant bool) (*domain.Person, error)
}
type StatLogic interface {
	StatsGetMutantVsHuman() (mutantCount, humanCount int64, ratio float32, err error)
}
type HandlerConstructor struct {
	Logger  Logger
	Storage Storage
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.Storage == nil {
		log.Fatal("missing Storage")
	}
	return interactor{
		logger:  c.Logger,
		storage: c.Storage,
	}
}
