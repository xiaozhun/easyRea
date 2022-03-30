cd %~dp0
CGO_ENABLED=0
set GOARCH=amd64
set GOOS=windows
go build -v -o easyReq.exe main.go
pause