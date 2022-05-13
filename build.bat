@echo off

set DIR_PROJECT=%~dp0
set PATH_RELEASE=%DIR_PROJECT%bin\CommandGo.exe

go build -ldflags "-s -w" -o %PATH_RELEASE% %DIR_PROJECT%\cmd\main.go
upx -9 %PATH_RELEASE%
ls -lh %PATH_RELEASE%