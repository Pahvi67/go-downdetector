#!/bin/sh

# get env variables if not inside docker
if ! [ -f /.dockerenv ]; then
    export $(grep -v '^#' .env | xargs)
    
    if [ "$ENVIRONMENT" = 'development' ]; then
        echo "Starting in dev env..."
        exec go run .
    else
        echo "Building and starting..."
        go build -o downDetector .
        exec ./downDetector
    fi
else
    exec ./main
fi