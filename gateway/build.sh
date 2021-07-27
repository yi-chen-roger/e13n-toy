#!/bin/bash
go mod download
go mod vendor
docker build . -t e13n_gateway