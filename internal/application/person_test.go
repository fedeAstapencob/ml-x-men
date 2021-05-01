package application_test

import (
	"github.com/golang/mock/gomock"
	"ml-x-men/internal/adapters/application.mock"
	"ml-x-men/internal/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonGetByDna_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(true)
	i := mock.NewMockedInteractor(mockCtrl)
	i.Storage.EXPECT().PersonGetByDna(person.Dna).Return(&person, nil).Times(1)

	retPerson, err := i.GetApplicationHandler().GetByDna(person.Dna)
	assert.NoError(t, err)
	assert.Equal(t, person, *retPerson)
}
