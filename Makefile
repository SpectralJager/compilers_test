build:
	go build -o grim cmd/interpreter/interpreter.go
run_test: build
	./grim examples/test.grim