package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var quantity int32
	quantity = 3

	fmt.Println(quantity)

	var length, height, width int32 = 1010, 2010, 3100

	length, height, width = 100, 200, 300
	fmt.Println(length, height, width)

	//fmt.Println("Hello, 世界", "ffff")
	//fmt.Println("Hello, 世界", "ffff")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("ffff"))

	fmt.Println('A')
	fmt.Println('爱')

	fmt.Println(10 / 3)

	fmt.Println(reflect.TypeOf("fd"))
	fmt.Println(reflect.TypeOf(1.3332))
	of := reflect.TypeOf(true)
	fmt.Println(of)

	var empty string

	fmt.Println(empty)
	fmt.Println(len(empty))

	size := 10
	size = 22
	fmt.Println(size)

	var num = 10
	var price = 43.12
	var total = float64(num) * price
	fmt.Println("Total is ", total)
	fmt.Println("Total is ", total)
	fmt.Println("Total is ", total)

	Hello()
}

func Hello() {
	fmt.Println("hello...")
}
