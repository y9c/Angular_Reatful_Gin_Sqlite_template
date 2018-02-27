GOCMD=go
BINARY_NAME=demoapi

# build Angular static file
#cd client
#npm install
#ng build --aot
#cd ..

## start API server
#go run server/main.go

## start Caddy
#ulimit -n 8192
#caddy -conf ./Caddyfile

ng build --prod --build-optimizer


run_%: %
	ulimit -n 8192 && ./$^


deps:
	$(GOCMD) get github.com/gin-gonic/contrib/static
	$(GOCMD) get github.com/gin-gonic/gin
	$(GOCMD) get github.com/jinzhu/gorm
	$(GOCMD) get github.com/jinzhu/gorm/dialects/sqlite

build:
	$(GOCMD) build -o $(BINARY_NAME) -v

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)

dev:
	$(GOCMD) build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

work: build

all: test build
