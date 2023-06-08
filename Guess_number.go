package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a program for guessing number game!

	number_system := rand.Intn(50) + 1
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess a number between 1 to 50: ")
	for true {
		val, err1 := reader.ReadString('\n')
		if err1 == nil {
			val = strings.TrimSpace(val)
			number_guessed, err2 := strconv.Atoi(val)
			if err2 == nil {
				if number_guessed == number_system {
					fmt.Printf("Correct guess! the number is %d", number_system)
					break
				} else if number_guessed > number_system {
					fmt.Print("Guess something lower: ")
				} else {
					fmt.Print("Guess something higher: ")
				}
			} else {
				log.Fatal(err2)
			}
		} else {
			log.Fatal(err1)
		}
	}
}