#!/bin/env sh

go install github.com/envoyproxy/protoc-gen-validate@v0.5.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0

rm -rf gen/go
mkdir -p gen/go

protoc \
  -I proto \
  -I "${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.5.1" \
  --go_out="paths=source_relative,Mtestsvc/person/v1/person.proto=github.com/sgtsquiggs/validate-package-name-error/gen/go/testsvc/person/v1/person;testsvc_person_v1,Mtestsvc/test/v1/enums.proto=github.com/sgtsquiggs/validate-package-name-error/gen/go/testsvc/test/v1/person;testsvc_test_v1:gen/go" \
  --validate_out="lang=go,paths=source_relative:gen/go" \
  testsvc/person/v1/person.proto testsvc/test/v1/enums.proto

