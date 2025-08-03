package main

import "fmt"

func Hello(name string) string {
	const defaultName = "Hello, "

	if name == "" {
		name = "World"
	}

	return defaultName + name
}

func main() {
	fmt.Println(Hello("Chris"))

}
