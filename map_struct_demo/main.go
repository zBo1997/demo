// 每个go文件都需要归属于一个包
package main

import (
	"fmt"

	"github.com/go-viper/mapstructure/v2"
)

type Family struct {
	LastName   string
	FamilyName string
}
type Location struct {
	City string
}
type Person struct {
	Family    `mapstructure:",squash"`
	Location  `mapstructure:",squash"`
	FirstName string
}
type Animal struct {
	Name string
}

func main() {
	input := map[string]interface{}{
		"FirstName": "Mitchell",
		"LastName":  "Hashimoto",
		"City":      "San Francisco",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.FirstName)
	fmt.Println(result.LastName)
	fmt.Println(result.City)
}
