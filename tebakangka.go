package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var maximum, level, guess, counter int
	fmt.Printf("\033[2JSelect level:\n") // escape is ANSI clear screen
	for level := 1; level != 13; level++ {
		maximum := level * 200
		fmt.Printf("Level %d: 1 - %d\n", level, maximum)
	}
	time.Sleep(1 * time.Second)
	fmt.Print("Level: ")
	_, err := fmt.Scan(&level)
	if err != nil {
		panic(err)
	}
	maximum = level * 200
	number := rand.Intn(maximum)
	fmt.Printf("\n\n\n")
	for {
		fmt.Printf("Guess the number (1 - %d)! ", maximum)
		_, err := fmt.Scan(&guess)
		if err != nil {
			panic(err)
		}
		counter++
		switch true {
		case guess == number:
			fmt.Printf("Congrats, You got it in %d tries!\n", counter)
			time.Sleep(1400 * time.Millisecond)
		case guess < number:
			fmt.Println("Wrong. guess a bigger number.")
		case guess > number:
			fmt.Println("Wrong. Guess a smaller number.")
		default:
			fmt.Println("Input Invalid.")
		}
	}
}
