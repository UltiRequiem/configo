# Configo

[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/configo.svg)](https://pkg.go.dev/github.com/UltiRequiem/configo)

Dead Simple Config Management, load and persist config without having to think
about where and how.

## Install

```sh
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

If you already ran the code snippet from before, then this will work 👇

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

## Documentation

Is hosted on [pkg.go.dev](https://pkg.go.dev/github.com/UltiRequiem/configo).

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to start the repo ⭐

## Licence

Licensed under the MIT License 📄
