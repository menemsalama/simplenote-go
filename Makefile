GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/simplenote-go

APP_ENTRY=cmd/simplenote-go/main.go
APP_API_DIR=api

# installtion command for swag, if It's not installed
ifeq ($(shell which swag),)
	SWAG_INSTALL=go get -u github.com/swaggo/swag/cmd/swag
endif

$(DOCKER_CMD): clean docsgen
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) $(APP_ENTRY)

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

docsgen:
	$(SWAG_INSTALL)
	swag init -g $(APP_API_DIR)/api.go -o $(APP_API_DIR)/docs

run: docsgen
	go run $(APP_ENTRY)
