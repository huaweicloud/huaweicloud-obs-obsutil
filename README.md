Version 5.1.9

Dependent SDK Version: 
Go SDK 2.2.12


Install command:

ARM:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil

Linux:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil


Windows:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil

MacOs:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil





