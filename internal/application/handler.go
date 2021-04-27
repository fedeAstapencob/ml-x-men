package application

import "log"

type Handler interface {
	MutantLogic
}

type MutantLogic interface {
	IsMutant(matrix [][]byte) (bool, error)
}

type HandlerConstructor struct {
	Logger     Logger
	MutantRepo MutantRepository
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	//if c.MutantRepo == nil {
	//	log.Fatal("missing MutantRepo")
	//}
	return interactor{
		logger:     c.Logger,
		mutantRepo: c.MutantRepo,
	}
}
