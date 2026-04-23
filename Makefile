.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

fanin: vet
	go run ./cmd/fanin

workerpool: vet
	go run ./cmd/workerpool