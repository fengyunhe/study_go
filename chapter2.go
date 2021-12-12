package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var now = time.Now()
	println(now.Year())
	broken := "hello #@$342342(#*($&(*@"
	replacer := strings.NewReplacer("#", "o", "3", "7")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	grade, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	grade = strings.TrimSpace(grade)
	fmt.Println(grade)

	s := "\t formerly surrounded by space \n"
	fmt.Println(strings.TrimSpace(s))

	parseBool, _ := strconv.ParseBool("  ")
	fmt.Println(parseBool)

	fmt.Println("Please input float number:")
	numberReader := bufio.NewReader(os.Stdin)

	readString, err := numberReader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(readString), 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("number is ", num)

	var status string
	if num > 10 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
