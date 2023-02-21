# Example

```go
package main

import (
	"github.com/semichkin-gopkg/conv"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	boris := User{"Boris", 54}

	parsedByJson, _ := conv.Struct[map[string]any](boris)
	log.Println(parsedByJson) // map[age:54 name:Boris]
}
```