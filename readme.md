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

### [`Configo.Get`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L15)

### [`Configo.GetAll`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L19)

### [`Configo.Set`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L23)

### [`Configo.SetAll`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L27)

### [`Configo.Has`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L31)

### [`Configo.Delete`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L39)

### [`Configo.Size`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L43)

### [`Configo.Clear`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L47)

### [`Configo.Path`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L51)

### [`Configo.Save`](https://github.com/UltiRequiem/configo/blob/main/configo.go#L55)

## Licence

This project is licensed under the [MIT Licence](./license).
