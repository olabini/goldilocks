#!/usr/bin/env bash

DIR=`pwd`

echo $PATH
sudo -E -i $DIR/run_tests_in.sh $DIR
