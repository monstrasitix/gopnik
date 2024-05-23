all:
	go run ./cmd/main.go

dev:
	air -c .air.toml ./cmd/main.go
