package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Olá, " + name + "!"
}

func main() {
	fmt.Println(Hello("Yago"))
}
