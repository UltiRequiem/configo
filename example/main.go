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

	config.Set("awesome", 777)

	fmt.Println(config.Get("awesome")) //=> 777

	fmt.Println(config.Has("awesome")) //=> true

	config.Delete("awesome")

	fmt.Println(config.Get("awesome")) //=> nil

	fmt.Println(config.Size()) //=> 0

	config.Set("foo", "bar")

	fmt.Println(config.Size()) //=> 1

	// If you want persistence of data
	errorSaving := config.Save()

	if errorSaving != nil {
		log.Fatal(errorSaving)
	}
}
