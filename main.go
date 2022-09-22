package main

import (
	"flag"
	"log"

	"hezzlTestTask/internal"
)

func main() {
	var pathConfig string
	flag.StringVar(&pathConfig, "config", "config/config.json", "")
	flag.Parse()

	config, err := internal.Parse(pathConfig)
	if err != nil {
		log.Fatalln("config not found")
	}

	internal.Start(config)
}
