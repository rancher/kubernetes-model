#!/bin/bash
set -e

cd $(dirname $0)/../generator

curl http://localhost:8080/swaggerapi/api/v1 > schemas.json
echo Saved schemas.json

echo -n Generating go code...
go run generator.go
echo " Done"

gofmt -w ../model/generated_*
echo Formatted code

echo Success

