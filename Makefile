.PHONY: all build run check help mocks clean test docs

BIN=cypt

all: lint check build

build:
	go build -o "${BIN}"

test:
	go test ./...

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	golangci-lint run -v

check:
	go fmt ./...
	go vet ./...

run: build
	"./${BIN}"

clean:
	go clean

mocks:
	mockery --all --dir internal/dddcore/ --output test/mocks/dddcore
	mockery --all --dir internal/user/ --output test/mocks/user
	mockery --all --dir internal/logger/ --output test/mocks/logger

upgrade:
	go get -u ./...

docs:
	swag init -o docs/swagger

help:
	echo "make         檢查後編譯"
	echo "make build   編譯"
	echo "make test    跑測試"
	echo "make check   檢查格式"
	echo "make run     直接執行"
	echo "make mocks   產生測試用的mock"
	echo "make upgrade 升級所有套件"
	echo "make docs    產生API文件"
