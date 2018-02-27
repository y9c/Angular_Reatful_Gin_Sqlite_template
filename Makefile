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

run_%: %
	ulimit -n 8192 && ./$^

#all: run_progname
