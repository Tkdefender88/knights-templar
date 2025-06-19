from golang:alpine as builder

workdir /src
copy go.mod go.sum ./
run go mod download
add ./assets assets
add ./cmd ./cmd
add ./internal ./internal
run go build -o ./bin/main ./cmd/main.go
expose 42069
entrypoint ["./bin/main"]
