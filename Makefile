default: build test

clean:
	@rm lib/*

build:
	@echo "Compiling Go source..."
	@go build -o lib/libgocrypt.so -buildmode=c-shared gocrypt.go
	@echo "Done!"
test:
	@echo "Running Go tests..."
	@go test -v -cover
	@echo "Running Python tests..."
	@pytest -s
