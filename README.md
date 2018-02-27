# Angular Reatful Gin Sqlite template

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

> Go

```bash
go get github.com/gin-gonic/contrib/static
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
```

## Run

> Compile and Run project

```bash
# build Angular static file
cd client
npm install
ng build --aot
cd ..

## start API server
go run server/main.go

## start Caddy
ulimit -n 8192
caddy -conf ./Caddyfile
```

Open `http://0.0.0.0:8888/`

## Changelog

- [x] use gin to route both static page and API
- [x] split router in differnt files ([ref](https://stackoverflow.com/questions/47115731/how-to-split-my-resources-into-multiply-files))
- [x] expose the API to the client
- [x] use caddy Proxy for /index.html and /api

