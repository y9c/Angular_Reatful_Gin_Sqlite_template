NGCMD=ng
GOCMD=go
CADDYCMD=caddy

API_BINARY=demoapi

TIMESTAMP:=$(shell date +%Y-%m-%d\ %H:%M:%S)
NODE_VERSION:=$(shell node -v)
NPM_VERSION:=$(shell npm -v)
GO_VERSION:=$(shell go version)

.PHONY: show build_go build_ng test clean

show:
	@ echo Timestamp: "$(TIMESTAMP)"
	@ echo Node Version: $(NODE_VERSION)
	@ echo npm_version: $(NPM_VERSION)
	@ echo go_version: $(GO_VERSION)

deps_global:
	curl https://getcaddy.com | bash -s http.git,http.ratelimit
	npm install --global @angular/cli

deps_ng:
	cd ./client; \
		npm install

deps_go:
	$(GOCMD) get github.com/gin-gonic/contrib/static
	$(GOCMD) get github.com/gin-gonic/gin
	$(GOCMD) get github.com/jinzhu/gorm
	$(GOCMD) get github.com/jinzhu/gorm/dialects/sqlite

deps: deps_global deps_ng deps_go

build_ng: server
	cd ./client; \
		$(NGCMD) build --prod --build-optimizer

build_go: server
	cd ./server; \
		$(GOCMD) build -ldflags="-s -w" -o $(API_BINARY) main.go; \
		upx --brute $(API_BINARY)

dev:
	caddy -conf ./server/caddy/Caddyfile_dev

work: build_ng build_go
	./server/$(BINARY_NAME)
	$(CADDYCMD) -conf ./server/caddy/Caddyfile

clean:
	rm -rf ./server/dist
	rm -rf ./server/$(API_BINARY)

all: deps work
	@ echo "Visit http://0.0.0.0:8888"
