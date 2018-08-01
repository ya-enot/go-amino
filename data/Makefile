.PHONY: docs
REPO:=github.com/ya-enot/go-wire/data

docs:
	@go get github.com/davecheney/godoc2md
	godoc2md $(REPO) > README.md

test:
	go test ./...
