package main

import (
	"flag"
	"log"

	"hezzlTestTask/config"
	"hezzlTestTask/internal/app"
)

func main() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "config/config.json", "")
	flag.Parse()

	if config.Parse(pathConfig) != nil {
		log.Fatalln("config not found")
	}

	app.Start(&config.Config)
}
