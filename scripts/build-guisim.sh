#!/bin/bash

# notice how we avoid spaces in $now to avoid quotation hell in go build command
sudo apt-get install gcc-multilib
sudo apt-get install gcc-mingw-w64
now=$(date --utc +%FT%T%Z)
GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
  go build -o guisim.exe -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" ./cmd/guisim/ 

# ls