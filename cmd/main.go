package main

import (
	"cypt/internal/registry"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("configs/.env")

	reg := registry.NewRegistry()
	app := reg.NewAppController()

	output, err := app.User.Rename("f7e41e07-c9cf-47bd-972f-64fec0882f20", "chuck10")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output.Result, output.Ret.Username)
	time.Sleep(time.Second)
}
