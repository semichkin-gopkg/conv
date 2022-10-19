# Example

```go
package main

import (
	"github.com/semichkin-gopkg/converter"
	"log"
)

type User struct {
	Name string `json:"name" some:"Name"`
	Age  uint   `json:"age" some:"Age"`
}

func main() {
	boris := User{"Boris", 54}

	parsedByJson, _ := converter.Convert[map[string]any](boris)
	log.Println(parsedByJson) // map[age:54 name:Boris]

	parsedBySome, _ := converter.Convert[map[string]any](boris, func(c *converter.Params) {
		c.Tag = "some"
	})
	log.Println(parsedBySome) // map[Age:54 Name:Boris]
}
```