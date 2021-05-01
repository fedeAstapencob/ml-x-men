package json_formatter

import "ml-x-men/internal/domain"

type PersonResp struct {
	ID       uint   `json:"id"`
	Dna      string `json:"dna"`
	IsMutant bool   `json:"isMutant"`
}

func NewPersonResp(person domain.Person) PersonResp {
	return PersonResp{
		ID:       person.ID,
		Dna:      person.Dna,
		IsMutant: person.IsMutant,
	}
}
