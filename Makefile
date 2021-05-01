all:
	go build -o ./cmd/bin/redis.exe ./cmd/cli/*.go
dev:
	go run ./cmd/cli/*.go
run:
	go build -o ./cmd/bin/redis.exe ./cmd/cli/*.go && ./cmd/bin/redis.exe
