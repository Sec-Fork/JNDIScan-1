#!/bin/bash

workdir=$(cd $(dirname $0); pwd)
echo $workdir
if [ -f "$workdir/JNDIScan-Linux" ]
  then
    rm "$workdir/JNDIScan-Linux"
fi

if [ -f "$workdir/JNDIScan-Darwin" ]
  then
    rm "$workdir/JNDIScan-Darwin"
fi

if [ -f "$workdir/JNDIScan-Windows.exe" ]
  then
    rm "$workdir/JNDIScan-Windows.exe"
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o JNDIScan-Linux
echo start linux success

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o JNDIScan-Darwin
echo start darwin success

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o JNDIScan-Windows.exe
echo start windows success