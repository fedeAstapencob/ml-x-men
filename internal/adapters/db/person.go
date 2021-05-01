package db

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"ml-x-men/internal/domain"
)

type PersonDB struct {
	gorm.Model
	Dna      string
	IsMutant sql.NullBool
}

func (PersonDB) TableName() string {
	return "person"
}

const (
	selectByDnaSQL = `SELECT id, dna, is_mutant FROM person WHERE dna = ?`
)

func (db DB) PersonCreate(dna string, isMutant bool) (*domain.Person, error) {
	var personDB = &PersonDB{Dna: dna}
	personDB.IsMutant = sql.NullBool{
		Bool:  isMutant,
		Valid: true,
	}
	result := db.Create(personDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.Person{
		ID:       personDB.ID,
		Dna:      dna,
		IsMutant: isMutant,
	}, nil
}

func (db DB) PersonGetByDna(dna string) (*domain.Person, error) {
	var person domain.Person
	var personDB PersonDB
	result := db.First(&personDB, "dna = ?", dna)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//person not found
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	person = domain.Person{
		ID:       personDB.ID,
		Dna:      personDB.Dna,
		IsMutant: personDB.IsMutant.Bool,
	}
	return &person, nil
}
