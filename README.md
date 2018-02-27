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

> Go

```bash
go get github.com/gin-gonic/contrib/static
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
```

## Compile and Run

Tested on linux machine only!

> Run on dev machine

Compile and Run project

```bash
make dev
```

> Run on work machine

```bash
make run
```

Open `http://0.0.0.0:8888/`

## Changelog

- [x] use gin to route both static page and API
- [x] split router in differnt files ([ref](https://stackoverflow.com/questions/47115731/how-to-split-my-resources-into-multiply-files))
- [x] expose the API to the client
- [x] use caddy Proxy for /index.html and /api
- [ ] write a makefile

