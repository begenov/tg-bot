package main

import (
	"fmt"
	"log"

	"github.com/begenov/tg-bot/internal/app"
	"github.com/begenov/tg-bot/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(cfg, err)
		return
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalln(err, "run function")
		return
	}
}
