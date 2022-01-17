@echo off
echo start build
if exist %cd%\JNDIScan-Linux (
    del JNDIScan-Linux
)
if exist %cd%\JNDIScan-Darwin% (
    del JNDIScan-Darwin
)
if exist %cd%\JNDIScan-Windows.exe (
    del JNDIScan-Windows.exe
)

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o JNDIScan-Linux
echo start linux success

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o JNDIScan-Darwin
echo start darwin success

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o JNDIScan-Windows.exe
echo start windows success
