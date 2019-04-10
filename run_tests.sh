#!/usr/bin/env bash

DIR=`pwd`

sudo -E -i $DIR/run_tests_in.sh $DIR $GOPATH $GOROOT
