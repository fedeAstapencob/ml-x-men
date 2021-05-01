mock:
	$(GOPATH)/bin/mockgen -source=./internal/application/interactor.go -destination=./internal/adapters/application.mock/interactor.go -package=mock && \
    $(GOPATH)/bin/mockgen -source=./internal/application/handler.go -destination=./internal/adapters/application.mock/handler.go -package=mock
