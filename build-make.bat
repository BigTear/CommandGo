@echo off
set GOOS=windows
set GOARCH=386
go.exe build -ldflags "-s -w" -o make.exe make.go
tools\upx.exe -9 -q make.exe
echo Done
echo Just run make.
