# Configo

[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/configo.svg)](https://pkg.go.dev/github.com/UltiRequiem/configo)

Dead Simple Config Management, easily load and persist config without having to think about where and how.

## Install

```
go get github.com/UltiRequiem/configo
```

## Usage

```go
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
```

## API

### `NewConfig`

### `Configo.Get`

### `Configo.GetAll`

### `Configo.Set`

### `Configo.SetAll`

### `Configo.Has`

### `Configo.Size`

### `Configo.Clear`

### `Configo.Path`

## Licence

This project is licensed under the [MIT Licence](./license).
