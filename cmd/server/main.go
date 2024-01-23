package main

import (
	"fmt"
	"go-products/configs"
)

func main() {
	config, _ := configs.Load(".")
	fmt.Println(config.DBDriver)
}
