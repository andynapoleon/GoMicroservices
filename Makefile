# Makefile

# Target for compiling .proto files
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.

# Target for running Go tests with race detection
test:
	go test -race ./...
