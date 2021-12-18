package main

import (
	"fmt"
	"go.firegod.cn/greetings"
)

func main() {
	message := greetings.Hello("Yang yan")
	fmt.Println(message)
}
