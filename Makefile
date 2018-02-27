# color
Color_Off='\033[0m'       # Text Reset
BBlack='\033[1;30m'       # Bold Black
BRed='\033[1;31m'         # Bold Red
BGreen='\033[1;32m'       # Bold Green
BYellow='\033[1;33m'      # Bold Yellow
BBlue='\033[1;34m'        # Bold Blue
BPurple='\033[1;35m'      # Bold Purple
BCyan='\033[1;36m'        # Bold Cyan
BWhite='\033[1;37m'       # Bold White

NG_CMD=ng
GO_CMD=go
CADDY_CMD=caddy
UPX_CMD=upx

API_BINARY=demoapi

TIMESTAMP:=$(shell date +%Y-%m-%d\ %H:%M:%S)
NODE_VERSION:=$(shell node -v)
NPM_VERSION:=$(shell npm -v)
GO_VERSION:=$(shell go version)

.PHONY: show build_go build_ng run_dev test clean

show:
	@ echo -e Timestamp: ${BCyan}$(TIMESTAMP)${Color_Off}
	@ echo -e Node Version: ${BCyan}$(NODE_VERSION)${Color_Off}
	@ echo -e npm_version: ${BCyan}$(NPM_VERSION)${Color_Off}
	@ echo -e go_version: ${BCyan}$(GO_VERSION)${Color_Off}

deps_global:
	@ #echo -e ${BYellow}Installing Caddy Server...${Color_Off}
	@ #curl https://getcaddy.com | bash -s personal http.git,http.ratelimit
	@ #echo -e ${BYellow}Installing Angular CLI Tools...${Color_Off}
	@ #npm install --global @angular/cli

deps_ng:
	@ echo -e ${BYellow}Installing Node package...${Color_Off}
	@ cd ./client; \
		npm install

deps_go:
	@ echo -e ${BYellow}Installing Go package...${Color_Off}
	@ $(GO_CMD) get github.com/gin-gonic/contrib/static
	@ $(GO_CMD) get github.com/gin-gonic/gin
	@ $(GO_CMD) get github.com/jinzhu/gorm
	@ $(GO_CMD) get github.com/jinzhu/gorm/dialects/sqlite

deps: deps_global deps_ng deps_go

run_dev: deps
	@ echo -e ${BBlue}Runing in dev mode...${Color_Off}
	@ cd ./client; \
		$(NG_CMD) server > ../server/log/ng.dev.log &
	@ cd ./server; \
		$(GO_CMD) run main.go > ./log/gin.dev.log &
	@ $(CADDY_CMD) -conf ./server/caddy/Caddyfile_dev > ./server/log/caddy.run.log &

build_ng: server
	@ echo -e ${BYellow}Building client...${Color_Off}
	@ cd ./client; \
		$(NG_CMD) build --prod --build-optimizer

build_go: server
	@ echo -e ${BYellow}Building server/API...${Color_Off}
	@ cd ./server; \
		$(GO_CMD) build -ldflags="-s -w" -o $(API_BINARY) main.go; \
		$(UPX_CMD) --brute $(API_BINARY)

work: deps build_ng build_go
	./server/$(BINARY_NAME)
	$(CADDY_CMD) -conf ./server/caddy/Caddyfile

clean:
	@ rm -rf ./server/log/*.log
	@ rm -rf ./server/dist
	@ rm -rf ./server/$(API_BINARY)

all: deps work
	@ echo "Visit http://0.0.0.0:8888"
