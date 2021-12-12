package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 2

	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it ?")
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	success := false
	for x := 0; x <= 10; x++ {
		fmt.Println("You have", 10-x, "guesses left.")
		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. your guess was low.")
		} else if guess > target {
			fmt.Println("Oops, your guess was high")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
