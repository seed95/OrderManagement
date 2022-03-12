package main

import (
	"Product/internal"
	"Product/internal/repo/product"
	nativeLog "log"
	"os"
)

func main() {

	configPrefix := os.Getenv("CONFIG_PREFIX")
	config := internal.NewConfig(configPrefix)

	productRepo, err := product.New(&product.Setting{Config: &config.ProductRepo})
	if err != nil {
		nativeLog.Fatal(err)
	}

	_ = productRepo

}
