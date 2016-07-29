.PHONY: all
default: install

lib:
	go install

cmd: lib
	go get github.com/kevinbuch/pong/cmd/pong

install: cmd

test: lib
	@cd test && go test -v
