test: lint vet
	go test

vet:
	go vet

lint:
	$(shell go list -f {{.Target}} golang.org/x/lint/golint) .

run: test
	go run cmd/detect-and-shoot/main.go 

demo:
	go build cmd/detect-and-shoot/demo.go
	./demo

.PHONY: test, vet, lint, run, demo
