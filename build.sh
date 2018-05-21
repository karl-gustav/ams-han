#!/bin/bash

set -e

if [ -d dist/ ] ; then
    rm -rf dist/
fi

mkdir dist/

GOOS=linux GOARCH=arm GOARM=6 go build -o mbus.reader
