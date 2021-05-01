package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "ml-x-men/internal/adapters/application.mock"
	"testing"
)

func TestInteractor_StatsGetMutantVsHuman_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	i := mock.NewMockedInteractor(mockCtrl)
	i.Storage.EXPECT().StatsGetMutantVsHuman().Return(int64(1), int64(4), float32(0.75), nil).Times(1)
	mutantCount, humanCount, ratio, err := i.GetApplicationHandler().StatsGetMutantVsHuman()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), mutantCount)
	assert.Equal(t, int64(4), humanCount)
	assert.Equal(t, float32(0.75), ratio)
}
