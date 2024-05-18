#!/usr/bin/env bash

protoc --proto_path=. --proto_path=./third_party  \
       --go_out=paths=source_relative:.  \
       --go-errors_out=paths=source_relative:. tools/protoc-gen-go-errors-code/errors/errors.proto
