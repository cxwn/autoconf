SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o .\bin\autoconf-linux-amd64 autoconf.go
