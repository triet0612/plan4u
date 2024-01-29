migrate:
	rm ./build/* || true
	go run ./cmd/migrate/migrate.go
	cp -r ./view/ ./build/
build_app:
	go build -o ./build/ ./cmd/server/main.go
run:
	make build_app && cd ./build && ./main
debug:
	make migrate && make build_app && cd ./build && PORT=8000 ./main
