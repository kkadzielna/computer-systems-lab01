package main

import "fmt"
import "time"

func main() {
	for i := 0; i < 20; i++ {
		go fmt.Println("Hello World from routine", i)
	}
	time.Sleep(1 * time.Second)
}