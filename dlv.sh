#!/bin/sh
PATH_TO_PROGRAM=$1
cd $PATH_TO_PROGRAM
dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2
