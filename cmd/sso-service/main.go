package main

import (
	"fmt"
	"sso-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
