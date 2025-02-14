package main

import "fmt"
import "time"

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
}
