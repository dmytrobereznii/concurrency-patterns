.DEFAULT_GOAL := run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

fanin: vet
	go run ./cmd/fanin

workerpool: vet
	go run ./cmd/workerpool

workerpool2: vet
	go run ./cmd/workerpool2

semaphore: vet
	go run ./cmd/semaphore

semaphore-doc:
	go doc ./cmd/semaphore

graceful:
	go run ./cmd/gracefulshutdown