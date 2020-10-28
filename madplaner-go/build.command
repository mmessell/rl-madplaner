#!/usr/bin/env bash

BASEDIR=$(dirname "$0")

cd "$BASEDIR"

go build program.go config.go
