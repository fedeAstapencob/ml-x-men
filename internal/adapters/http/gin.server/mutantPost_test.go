package gin_server_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/h2non/baloo.v3"
	mock "ml-x-men/internal/adapters/application.mock"
	server "ml-x-men/internal/adapters/http/gin.server"
	"ml-x-men/internal/adapters/logger"
	"ml-x-men/internal/testData"
	"ml-x-men/internal/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var personPostPath = "/api/mutant"

func TestMutantPost_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(true)
	ucHandler := mock.NewMockHandler(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrixDna := utils.BuildMatrixDna(dnaArray)
	ucHandler.EXPECT().
		PersonCreate(person.Dna, person.IsMutant).
		Return(&person, nil).
		Times(1)
	ucHandler.EXPECT().
		GetByDna(person.Dna).
		Return(nil, nil).
		Times(1)
	ucHandler.EXPECT().
		IsMutant(matrixDna).
		Return(true, nil).
		Times(1)
	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
			"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusOK).
		Done()
}

func TestHumanPost_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(false)
	ucHandler := mock.NewMockHandler(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrixDna := utils.BuildMatrixDna(dnaArray)
	ucHandler.EXPECT().
		PersonCreate(person.Dna, person.IsMutant).
		Return(&person, nil).
		Times(1)
	ucHandler.EXPECT().
		GetByDna(person.Dna).
		Return(nil, nil).
		Times(1)
	ucHandler.EXPECT().
		IsMutant(matrixDna).
		Return(false, nil).
		Times(1)
	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
			"dna": ["ATGCGG","CAGTGC","TTATGT","AGACAG","GCGTCA","TCACTG"]
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusForbidden).
		Done()
}

func TestPersonPostBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := mock.NewMockHandler(mockCtrl)

	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusBadRequest).
		Done()
}

func TestPersonPostErrorGettingByDna(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(true)
	ucHandler := mock.NewMockHandler(mockCtrl)

	ucHandler.EXPECT().
		GetByDna(person.Dna).
		Return(nil, errors.New("error getting by dna")).
		Times(1)
	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
			"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusInternalServerError).
		Done()
}

func TestPersonPostErrorIsMutant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(false)
	ucHandler := mock.NewMockHandler(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrixDna := utils.BuildMatrixDna(dnaArray)
	ucHandler.EXPECT().
		GetByDna(person.Dna).
		Return(nil, nil).
		Times(1)
	ucHandler.EXPECT().
		IsMutant(matrixDna).
		Return(false, errors.New("Error evaluating IsMutant")).
		Times(1)
	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
			"dna": ["ATGCGG","CAGTGC","TTATGT","AGACAG","GCGTCA","TCACTG"]
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusInternalServerError).
		Done()
}

func TestPersonErrorPersonCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	person := testData.Person(false)
	ucHandler := mock.NewMockHandler(mockCtrl)
	dnaArray := strings.Split(person.Dna, ",")
	matrixDna := utils.BuildMatrixDna(dnaArray)
	ucHandler.EXPECT().
		PersonCreate(person.Dna, person.IsMutant).
		Return(nil, errors.New("Error creating person")).
		Times(1)
	ucHandler.EXPECT().
		GetByDna(person.Dna).
		Return(nil, nil).
		Times(1)
	ucHandler.EXPECT().
		IsMutant(matrixDna).
		Return(false, nil).
		Times(1)
	gE := gin.Default()
	router := server.NewRouter(ucHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	_ = baloo.New(ts.URL).
		Post(personPostPath).
		BodyString(`
		{
			"dna": ["ATGCGG","CAGTGC","TTATGT","AGACAG","GCGTCA","TCACTG"]
		}`).
		Expect(t).
		JSONSchema(testData.PersonRespDefinition).
		Status(http.StatusInternalServerError).
		Done()
}
