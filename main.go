package main

import (
	"fmt"
)

func Welcome(s string) string {
	return s
}

func main() {
	s := "Welcome to Zone01"
	fmt.Println(Welcome(s))
}
