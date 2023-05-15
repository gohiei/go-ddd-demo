# Domain-Driven Design + Clean Architecture for golang

## Steup & Run
* 請設定 configs/.env，可參考 configs/.env.example
* go run cmd/main.go

## Test
* Use mockery to generate mock files: `mockery --all --dir internal/user/ --output test/mocks/user`
* Run all tests: `go test ./...`
* Run some tests: `go test cypt/internal/user/usecase`

## Todo
* ~EventBus~
* Restful API (gin?)
* ~Database Read/Write Split (gorm?)~
* IoC/DI package (?)
* Tests
* Logger
