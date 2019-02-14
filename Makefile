install:
	dep ensure

test:
	go test ./... -v

.PHONY: install test
