package main

func main() {
	myMap := make(map[int]struct{})
	myMap[6] = struct{}{}
	myMap[2] = struct{}{}

	for k := range myMap {
		println(k)
	}

	mynotifyChan := make(map[string]chan int)
	mynotifyChan["test"] = make(chan int, 1)
	mynotifyChan["test"] <- 1

	println(<-mynotifyChan["test"])
}
