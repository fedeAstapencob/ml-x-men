BINARY=ml-x-men
BUILD=`git rev-parse HEAD`
COVERAGEALL=coverage-all.out
COVERHTML=test-coverage

test:
	go test ./...
cover:
	go test ./... -coverprofile=coverage.out

mock:
	$(GOPATH)/bin/mockgen -source=./internal/application/interactor.go -destination=./internal/adapters/application.mock/interactor.go -package=mock && \
    $(GOPATH)/bin/mockgen -source=./internal/application/handler.go -destination=./internal/adapters/application.mock/handler.go -package=mock
