.PHONY: all build run check help mocks clean test

BIN=cypt

all: check build

build:
	go build -o "${BIN}"

test:
	go test ./...

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

help:
	echo "make         檢查後編譯"
	echo "make build   編譯"
	echo "make test    跑測試"
	echo "make check   檢查格式"
	echo "make run     直接執行"
	echo "make mocks   產生測試用的mock"
	echo "make upgrade 升級所有套件"
