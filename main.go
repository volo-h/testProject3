package main

import "fmt"

func main() {
	msg := sayHello("Alice")
	fmt.Println(msg)
}

func sayHello(name string) string {
	// Измените это на "Hello %s" вместо "Hi %s".
	return fmt.Sprintf("Hello %s", name)
}
