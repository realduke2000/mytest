#!/bin/bash
cd srv
go build -o ../rolex-srv main.go
cd ..
docker build -t rolex-srv:latest .