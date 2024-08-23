 include .env
 include .env.local
 export $(shell sed 's/=.*//' .env)
 export $(shell sed 's/=.*//' .env.local)

run-cli:
	go run ./cli/main.go

run: 
	CompileDaemon -command="./main" -include="*.html"

build: 
	rm -rf ./dist
	mkdir ./dist
	mkdir ./dist/linux-amd64
	mkdir ./dist/darwin-amd64
	mkdir ./dist/windows-arm64
	mkdir ./dist/windows-amd64
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux-amd64 ./cmd/web/main.go
	GOOS=darwin GOARCH=amd64 go build -o ./dist/darwin-amd64 ./cmd/web/main.go
	GOOS=windows GOARCH=arm64 go build -o ./dist/windows-arm64 ./cmd/web/main.go
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows-amd64 ./cmd/web/main.go

generate:
	sqlc generate

migration:
	migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	migrate -path database/migrations -database "$(DB_DRIVER)://$(DB_DSN)" -verbose up

migrate-down:
	migrate -path database/migrations -database "$(DB_DRIVER)://$(DB_DSN)" -verbose down 1

analyze:
	gocloc .