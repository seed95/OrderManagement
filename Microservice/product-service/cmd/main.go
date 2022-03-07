package main

import (
	"Product/internal"
	"Product/internal/repo/carpet"
	nativeLog "log"
	"os"
)

func main() {

	configPrefix := os.Getenv("CONFIG_PREFIX")
	config := internal.NewConfig(configPrefix)

	productRepo, err := carpet.New(&carpet.Setting{Config: &config.ProductRepo})
	if err != nil {
		nativeLog.Fatal(err)
	}

	_ = productRepo

}
