package application

import "ml-x-men/internal/domain"

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger  Logger
	storage Storage
}

type Logger interface {
	Log(...interface{})
}

type Storage interface {
	PersonCreate(dna string, isMutant bool) (*domain.Person, error)
	PersonGetByDna(dna string) (*domain.Person, error)
	StatsGetMutantVsHuman() (mutantCount, humanCount int64, ratio float32, err error)
}
