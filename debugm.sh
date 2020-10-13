#!/usr/bin/env bash

go build -o fabric-first-go-app
dlv --listen=:2345 --headless=true --api-version=2 exec ./fabric-first-go-app api