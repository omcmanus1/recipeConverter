BINARY_NAME=build
.DEFAULT_GOAL := run

build: clean
	@GOARCH=arm64 GOOS=darwin go build -o ../bin/${BINARY_NAME}/darwin
	@GOARCH=amd64 GOOS=linux go build -o ../bin/${BINARY_NAME}/linux
	@GOARCH=amd64 GOOS=windows go build -o ../bin/${BINARY_NAME}/windows

run: build
	@../bin/${BINARY_NAME}/darwin

run_linux: build
	@../bin/${BINARY_NAME}/linux

run_windows: build
	@../bin/${BINARY_NAME}/windows

run_dev: build
	@ENVIRONMENT=local air

clean:
	@rm -rf ../bin/build

deploy:
	@../scripts/deploy_cloud_functions.sh

clean_logs:
	@rm -rf ../scripts/logs/deploy_cloud_functions/*

test: 
	go test -v ../... -count=1

test_coverage:
	go test ../... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet ../...
	shadow ../...

.PHONY:vet
