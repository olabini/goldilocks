#!/usr/bin/env bash

DIR=`pwd`

echo $PATH
echo $GOPATH
echo $GOROOT
go env
sudo -E -i $DIR/run_tests_in.sh $DIR
