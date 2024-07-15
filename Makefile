# include .env
# include .env.local
# export $(shell sed 's/=.*//' .env)
# export $(shell sed 's/=.*//' .env.local)

run: 
	CompileDaemon -command="./main" -include="*.html"

build: 
	rm -rf ./dist
	mkdir ./dist
	mkdir ./dist/linux-amd64
	mkdir ./dist/darwin-amd64
	mkdir ./dist/windows-arm64
	mkdir ./dist/windows-amd64
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux-amd64 ./main.go 
	GOOS=darwin GOARCH=amd64 go build -o ./dist/darwin-amd64 ./main.go
	GOOS=windows GOARCH=arm64 go build -o ./dist/windows-arm64 ./main.go 
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows-amd64 ./main.go 


analyze:
	gocloc .