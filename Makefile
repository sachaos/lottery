.PHONY: prepare test

prepare:
	go install github.com/golang/mock/mockgen
	go generate
	dep ensure -v

test: prepare
	go test -v
