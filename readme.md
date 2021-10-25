# Configo [![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/configo.svg)](https://pkg.go.dev/github.com/UltiRequiem/configo) [![CI](https://github.com/UltiRequiem/configo/workflows/CI/badge.svg)](https://github.com/UltiRequiem/configo/actions/workflows/ci.yaml)

Dead Simple Config Management, load and persist config without having to think about where and how.

## Install

```bash
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

	// Only if you want to persist values
	defer config.Save()

	config.Set("awesome", 777)

	fmt.Println(config.Get("awesome")) //=> 777

	fmt.Println(config.Has("awesome")) //=> true

	config.Delete("awesome")

	fmt.Println(config.Get("awesome")) //=> nil

	fmt.Println(config.Size()) //=> 0

	config.Set("foo", "bar")

	fmt.Println(config.Size()) //=> 1
}
```

If you already ran the code snippet from before, then this will work:

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

	fmt.Println(config.Get("foo")) // "bar"
}
```

## API

### [`NewConfig`](https://github.com/UltiRequiem/configo/blob/main/root.go#L9)

Returns a new instance of [`Configo`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L10).

### [`Configo.Get`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L15)

Returns the `value` on the `key` passed.

### [`Configo.GetAll`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L19)

Returns a map with all the config.

### [`Configo.Set`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L23)

Set the `value` in the `key` passed.

### [`Configo.SetAll`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L27)

Receives a `map[string]interface{}`.

### [`Configo.Has`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L31)

It tells you if it has the received `key`.

### [`Configo.Delete`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L39)

Deletes the `key` passed.

### [`Configo.Size`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L43)

Returns the quantity of the config properties.

### [`Configo.Clear`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L47)

Removes all the values of the config.

### [`Configo.Path`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L51)

Gives you the path were the config is stored.

### [`Configo.Save`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L55)

Saves the config for persistence.

## Licence

This project is licensed under the [MIT Licence](./license).
