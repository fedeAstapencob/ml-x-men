package personDB

import (
	"ml-x-men/internal/domain"
)
const (
	insertPersonSQL = `INSERT INTO person(dna, is_mutant) VALUES (?,?)`
	selectByDnaSQL = `SELECT * FROM person WHERE dna = ?`
)
func (db DB) Create(dna string, isMutant bool) (*domain.Person, error){

	person, err := db.Exec(insertPersonSQL, dna, isMutant)
	if err != nil {
		return nil, err
	}
	personId, err := person.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &domain.Person{
		ID:       personId,
		Dna:      dna,
		IsMutant: isMutant,
	}, nil
}

func (db DB) GetByDna(dna string) (*domain.Person, error){
	var person domain.Person
	row, err := db.Queryx(selectByDnaSQL,dna)
	if err != nil {
		return &person, err
	}
	err = row.Scan(person)
	return &person, err
}