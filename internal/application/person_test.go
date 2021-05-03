package application_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"ml-x-men/internal/adapters/application.mock"
	"ml-x-men/internal/testData"
	"ml-x-men/internal/utils"
	"strings"
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

func TestPersonGetByDna_errorCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(false)

	i := mock.NewMockedInteractor(mockCtrl)
	i.Storage.EXPECT().PersonGetByDna(person.Dna).Return(nil, fmt.Errorf("test error")).Times(1)

	retPerson, err := i.GetApplicationHandler().GetByDna(person.Dna)
	assert.Error(t, err)
	assert.Nil(t, retPerson)
}

func TestPersonCreate_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(false)
	i := mock.NewMockedInteractor(mockCtrl)
	i.Storage.EXPECT().PersonCreate(person.Dna, person.IsMutant).Return(&person, nil).Times(1)

	retPerson, err := i.GetApplicationHandler().PersonCreate(person.Dna, person.IsMutant)
	assert.NoError(t, err)
	assert.Equal(t, person, *retPerson)
}

func TestPersonGetPersonCreate_errorCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	person := testData.Person(false)

	i := mock.NewMockedInteractor(mockCtrl)
	i.Storage.EXPECT().PersonCreate(person.Dna, person.IsMutant).Return(nil, fmt.Errorf("test error")).Times(1)

	retPerson, err := i.GetApplicationHandler().PersonCreate(person.Dna, person.IsMutant)
	assert.Error(t, err)
	assert.Nil(t, retPerson)
}
func TestPersonIsMutant_true(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	person := testData.Person(true)

	i := mock.NewMockedInteractor(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrix := utils.BuildMatrixDna(dnaArray)
	isMutant, err := i.GetApplicationHandler().IsMutant(matrix)
	assert.NoError(t, err)
	assert.True(t, isMutant)
}

func TestPersonIsMutant_false(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	person := testData.Person(false)

	i := mock.NewMockedInteractor(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrix := utils.BuildMatrixDna(dnaArray)
	isMutant, err := i.GetApplicationHandler().IsMutant(matrix)
	assert.NoError(t, err)
	assert.False(t, isMutant)
}
