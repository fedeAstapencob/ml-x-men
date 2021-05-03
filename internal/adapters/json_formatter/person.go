package json_formatter

import "ml-x-men/internal/domain"

type PersonResp struct {
	Dna      string `json:"dna"`
	IsMutant bool   `json:"isMutant"`
}

func NewPersonResp(person domain.Person) PersonResp {
	return PersonResp{
		Dna:      person.Dna,
		IsMutant: person.IsMutant,
	}
}
