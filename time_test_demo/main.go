package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().After(time.Date(2027, 6, 1, 0, 0, 0, 0, time.UTC)))
}
