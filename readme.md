# passgen 🔐

Un generador de contraseñas seguras en Go.

### Instalación

```bash
go get github.com/Mike-droid/passgen
```

### Ejemplo de uso

```go
package main

import (
	"fmt"
	"github.com/Mike-droid/passgen"
)

func main() {
	cfg := passgen.Config {
		Length : 20,
		UseUpper : true,
		UseLower : true,
		UseNumbers : true,
		UseSymbols : true,
	}

	fmt.Println(passgen.Generate(cfg))
}

```

