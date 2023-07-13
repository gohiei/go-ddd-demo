# Domain-Driven Design + Clean Architecture for golang

## Steup & Run
* Create your own config file (refer to: configs/config.example.yaml)
* `go run main.go server -c <your config filename>`
  * More detail: `go run main.go --help`
  * More detail: `go run main.go server --help`

## Test
* Use mockery to generate mock files: `mockery --all --dir internal/user/ --output test/mocks/user`
* Run all tests: `go test ./...`
* Run some tests: `go test cypt/internal/user/usecase`
* See [installation](https://vektra.github.io/mockery/installation/)

## Release
* `make build`
* `./cypt server -c <config> -p 8080`

## Todo
* ~EventBus~
* ~Restful API (gin)~
* ~Database Read/Write Split (gorm)~
* ~IoC/DI package~ (Deprecated)
* Tests
* ~Logger~
* ~Customized Error~ ([ref](https://github.com/gohiei/go-ddd-demo/commit/11416ce5673785122497fe300e720a70e6831912))
* ~Configuration (viper)~
* ~Commandline tool (cobra)~
* ~Makefile~

## Recommendation
* [uber-go/goleak](https://github.com/uber-go/goleak)
* [uber-go/automaxprocs](https://github.com/uber-go/automaxprocs)
* [go-kit/kit](https://github.com/go-kit/kit)

## Modules

### Auth
* Use gin middleware for authentication
* Use jwt token for validation

### Logger
* Use gin middleware for event generation
* Use event for logging
