run: 
	go run ./cmd/grimlang/grimlang.go
cgc:
	GOGC=150 go run ./cmd/grimlang/grimlang.go