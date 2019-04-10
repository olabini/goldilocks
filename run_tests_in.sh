#!/usr/bin/env bash

echo "Running make test in $1"
cd $1
ulimit -a
make test
