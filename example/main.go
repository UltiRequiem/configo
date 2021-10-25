package main

import (
	"fmt"
	"log"

	"github.com/UltiRequiem/configo"
)

func main() {
	config, err := configo.NewConfig("example")

	if err != nil {
		log.Fatal(err)
	}

	config.Set("awesome", true)

	fmt.Println(config.Get("awesome")) //=> true

	config.Delete("awesome")

	fmt.Println(config.Get("awesome")) //=> nil

	fmt.Println(config.Size())
}
