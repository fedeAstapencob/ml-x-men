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
	"net/http"
	"net/http/httptest"
	"testing"
)

var tagsPath = "/api/stats"

func TestStatsGet_happyCase(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ucHandler := mock.NewMockHandler(mockCtrl)
		ucHandler.EXPECT().
			StatsGetMutantVsHuman().
			Return(int64(1), int64(1), float32(0.5), nil).
			Times(1)

		gE := gin.Default()
		router := server.NewRouter(ucHandler)
		router.Logger = logger.SimpleLogger{}
		router.SetRoutes(gE)
		ts := httptest.NewServer(gE)
		defer ts.Close()

		_ = baloo.New(ts.URL).
			Get(tagsPath).
			Expect(t).
			Status(http.StatusOK).
			JSONSchema(testData.StatsRespDefinition).
			Done()
	})
	t.Run("empty", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ucHandler := mock.NewMockHandler(mockCtrl)
		ucHandler.EXPECT().
			StatsGetMutantVsHuman().
			Return(int64(0), int64(0), float32(0), nil).
			Times(1)

		gE := gin.Default()
		router := server.NewRouter(ucHandler)
		router.Logger = logger.SimpleLogger{}
		router.SetRoutes(gE)
		ts := httptest.NewServer(gE)
		defer ts.Close()

		_ = baloo.New(ts.URL).
			Get(tagsPath).
			Expect(t).
			Status(http.StatusOK).
			JSONSchema(testData.StatsRespDefinition).
			Done()
	})

}

func TestStatsGet_fail(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ucHandler := mock.NewMockHandler(mockCtrl)
		ucHandler.EXPECT().
			StatsGetMutantVsHuman().
			Return(int64(0), int64(0), float32(0), errors.New("error")).
			Times(1)

		gE := gin.Default()
		router := server.NewRouter(ucHandler)
		router.Logger = logger.SimpleLogger{}
		router.SetRoutes(gE)
		ts := httptest.NewServer(gE)
		defer ts.Close()

		_ = baloo.New(ts.URL).
			Get(tagsPath).
			Expect(t).
			Status(http.StatusInternalServerError).
			JSONSchema(testData.StatsRespDefinition).
			Done()
	})

}
