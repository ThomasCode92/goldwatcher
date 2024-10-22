BINARY_NAME=GoldWatcher.app
APP_NAME=GoldWatcher
VERSION=1.0.0
BUILD_NO=1

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -icon icon.png -release
	rm -f go-for-gold

## run: builds and runs the application
run:
	env DB_PATH="./sqlite.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...
