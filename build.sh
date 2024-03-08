env CGO_ENABLED=0 go build -o bin/web goweb/*.go
env CGO_ENABLED=0 go build -o bin/user user/*.go