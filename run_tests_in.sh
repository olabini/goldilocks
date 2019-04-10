#!/usr/bin/env bash

cd $1
ulimit -l 16384

export GOPATH=$2
export GOROOT=$3
export PATH=$GOPATH/bin:$GOROOT/bin:/home/travis/bin:/home/travis/.local/bin:$PATH

make test
