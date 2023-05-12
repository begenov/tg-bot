package main

import (
	"fmt"

	"github.com/begenov/tg-bot/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(cfg, err)
		return
	}
}
