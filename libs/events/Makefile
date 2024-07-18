.PHONY: docs
REPO:=github.com/airchains-network/wasmbft/libs/events

docs:
	@go get github.com/davecheney/godoc2md
	godoc2md $(REPO) > README.md

test:
	go test -v ./...
