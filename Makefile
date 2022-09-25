LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=personhood-proof

export GO111MODULE=auto
GOENV:=GO111MODULE=auto

.PHONY: build
build:
	$(GOENV) go build -v -o $(LOCAL_BIN)/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)
	go build -v -o ./migrate ./cmd/migrate

.PHONY: run-local
run-local:
	make build && ./bin/$(PROJECT_NAME)

.PHONY: migrate
migrate:
	./migrate 