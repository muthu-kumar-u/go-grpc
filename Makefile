.PHONY: proto server

proto:
	buf dep update && buf generate

run-server:
	go run ./server

run-client:
	go run ./client
