#!/bin/bash

if [ -d "build" ]; then
    rm -rf build
fi

mkdir build

cp -r resources/* build/resources/

cd src

go build -ldflags "-s -w" -o "../build/GameServer.exe" main.go
