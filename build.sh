go build -ldflags="-s -w" -o kasir-api
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o kasir-api