// +build !netgo

package mock

import (
	"github.com/golang/mock/gomock"
	"log"
	"ml-x-men/internal/application"
)

// Interactor : is used in order to update its properties accordingly to each test conditions
type Interactor struct {
	Logger  *MockLogger
	Storage *MockStorage
}

type Tester struct {
	Calls      func(*Interactor)
	ShouldPass bool
}

type SimpleLogger struct{}

func (SimpleLogger) Log(logs ...interface{}) {
	log.Println(logs...)
}

//NewMockedInteractor : the Interactor constructor
func NewMockedInteractor(mockCtrl *gomock.Controller) Interactor {
	return Interactor{
		Logger:  NewMockLogger(mockCtrl),
		Storage: NewMockStorage(mockCtrl),
	}
}

//GetUCHandler : returns a application.interactor in order to call its methods aka the use cases to test
func (i Interactor) GetApplicationHandler() application.Handler {
	return application.HandlerConstructor{
		Logger:  i.Logger,
		Storage: i.Storage,
	}.New()
}
