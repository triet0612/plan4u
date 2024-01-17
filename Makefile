migrate:
	rm ../build/*.db || true
	go run ./cmd/migrate/migrate.go
build_app:
	go build -o ./build/ ./cmd/server/main.go
run:
	make build_app && cd ./build && ./main
debug:
	make build_app && cd ./build && PORT=8000 ./main
