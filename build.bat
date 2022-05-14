@echo off

set DIR_PROJECT=%~dp0
mkdir bin
set PATH_RELEASE=%DIR_PROJECT%bin\CommandGo.exe

set GOOS=windows
set GOARCH=386

go.exe build -ldflags "-s -w" -o %PATH_RELEASE% %DIR_PROJECT%\cmd\main.go
tools\upx.exe -q -9 %PATH_RELEASE%
ls -lh %PATH_RELEASE%
