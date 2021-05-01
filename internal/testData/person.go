package testData

import "ml-x-men/internal/domain"

var mutant = &domain.Person{
	ID:       1,
	Dna:      "ATGCGA,CAGTGC,TTATGT,AGAAGG,CCCCTA,TCACTG",
	IsMutant: true,
}
var human = &domain.Person{
	ID:       2,
	Dna:      "ATGCGG,CAGTGC,TTATGT,AGACAG,GCGTCA,TCACTG",
	IsMutant: false,
}

func Person(getMutant bool) domain.Person {
	if getMutant {
		return *mutant
	}
	return *human
}
