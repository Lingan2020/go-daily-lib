package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	user := `{"name":dj,age:18}`
	newValue, err := sjson.Set(user, "name", "dajun")
	fmt.Println(err, newValue)
}
