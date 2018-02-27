# `Angular + Reatful + Gin +  Sqlite` template

## Application Architecture

```diagram

    ----------
    | client | ←-----|
    ----------       |
       |             ↓
       |         ----------
       |         | server |
       |         ----------
       |render       ↑
       |         ---------
       ||------→ | Caddy |
       ||        ---------
       ↓|proxy       ↑
    --------     -------    ---------    --------
    | View |     | API | ←  | Model | ←  | Data |
    --------     -------    ---------    --------
        ↓            ↓          ↓            ↓
    -----------  -----------  -------  ----------
    | Angular |  | Restful |  | Gin |  | Sqlite |
    -----------  -----------  -------  ----------
        ↓                        ↓
    --------------           ----------
    | TypeScript |           | Golang |
    --------------           ----------

```

## Prerequisites

- **Client**:

> Node, npm

```bash
npm install --global @angular/cli
```

- **Server**:

> Go, caddy

```bash
go get github.com/gin-gonic/contrib/static
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
```

```bash
curl https://getcaddy.com | bash -s http.git,http.ratelimit
```

## Compile and Run

Tested on linux machine only!

> Run on development environment

```bash
# make dev

cd client
ng server &
cd ..
go run server/main.go &
caddy -conf ./server/caddy/Caddyfile_dev
```

> Run on production environment

```bash
# make run

cd client
ng build --prod
cd ../server
go build main.go
cd ..
./server/main
caddy -conf ./server/caddy/Caddyfile
```

Open `http://0.0.0.0:8888/`

## Changelog

- [x] use gin to route both static page and API
- [x] split router in differnt files ([ref](https://stackoverflow.com/questions/47115731/how-to-split-my-resources-into-multiply-files))
- [x] expose the API to the client
- [x] use caddy Proxy for /index.html and /api
- [ ] write a makefile
- [ ] Add Wiki page

## ISSUE

- Gin model is case insensitive, but Angular is case sensitive. (May raise error)

