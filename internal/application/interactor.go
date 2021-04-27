package application

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger     Logger
	mutantRepo MutantRepository
}

type MutantRepository interface {
}
