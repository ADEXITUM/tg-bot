.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot main.go

run: build
	./.bin/bot