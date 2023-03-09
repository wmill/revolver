
PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=./proto/user --go_opt=paths=source_relative \
    --go-grpc_out=./proto/user --go-grpc_opt=paths=source_relative \
    ./proto/user/user.proto