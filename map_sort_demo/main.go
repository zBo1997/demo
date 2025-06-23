// 每个go文件都需要归属于一个包
package main

import "sort"

func main() {
	data := make(map[string]int)
	data["apple"] = 3
	data["cherry"] = 5
	data["banana"] = 6

	keys := make([]string, len(data))
	j := 0
	for k := range data {
		keys[j] = k
		j++
	}

	sort.Strings(keys)
	for _, k := range keys {
		println(k, data[k])
	}
}
