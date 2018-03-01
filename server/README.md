# demo

server

## directory structure
.
├── main.go
├── models
│   ├── members.go
│   ├── papers.go
│   └── share.go
├── data
│   └── db.sqlite3
├── analysis
│   ├── manipulate_db.py
│   ├── Pipfile
│   └── Pipfile.lock
├── caddy
│   ├── Caddyfile
│   └── Caddyfile_dev
└── log
    ├── access.log
    ├── caddy.run.log
    ├── gin.dev.log
    └── ng.dev.log
