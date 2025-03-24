package main

import "fmt"

func main() {
	testData := make(map[string][]string, 10)
	testData["张三"] = []string{"1000", "1000"}
	testData["李四"] = []string{"1000", "1000", "18"}
	testData["王五"] = []string{"1000", "1000", "18"}
	for key, value := range testData {
		if key == "李四" {
			value = append(value, "1800")
			testData[key] = value
		}
	}
	fmt.Println(testData)
}
