build:
	go build -o grim cmd/interpreter/interpreter.go
run_test:
	./grim examples/test.grim