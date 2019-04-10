#!/usr/bin/env bash

echo "Running make test in $1"
cd $1
ulimit -l 16384

export GOPATH="/home/travis/gopath"
export GOROOT="/home/travis/.gimme/versions/go1.12.3.linux.amd64"
export PATH=$GOPATH/bin:$GOROOT/bin:/home/travis/bin:/home/travis/.local/bin:$PATH

make test
