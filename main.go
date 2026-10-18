package main

import (
	"bufio"
	"fmt"
	"os"
)

const petName string = "Momo"

var scanner = bufio.NewScanner(os.Stdin)
var petMemoriesRaw []string

func petSay(msg string) {
	fmt.Printf("%s: %s\n", petName, msg)
}
func petUserInputFeedback(fact string) {
	petSay("I'll try to remember that!")
	// petSay(fact)
	petMemoriesRaw = append(petMemoriesRaw, fact)
	// fmt.Printf("You have this much memories: %d\n", len(petMemoriesRaw))
}
func petMenu() {
	fmt.Println()
	petSay("What should we do? ")
	fmt.Println()
	fmt.Println("1. Teach me something")
	fmt.Println("2. What do you remember?")
	fmt.Println("3. Quit")
	var choice string
	fmt.Println()
	fmt.Printf("Choice[1/2/3]: ")
	if scanner.Scan() {
		choice = scanner.Text()
	}
	fmt.Println()
	switch choice {
	case "1":
		petSay("Learning time....")
		fact := inputUser()
		petUserInputFeedback(fact)
	case "2":
		if (len(petMemoriesRaw)) == 0 {
			petSay("I don't remember anything yet :( ,Teach me something")
		}
		petSay("Let me see what I remember!")
		for _, value := range petMemoriesRaw {
			petSay(value)
		}
	case "3":
		petSay("Byeee")
		os.Exit(0)
	default:
		petSay("I don't understand your choie ;( Retry plss....")
	}
}
func inputUser() string {
	var line string
	fmt.Printf("> ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading line:", err)
			os.Exit(1)
		}
		fmt.Println("\nNo Input Provided (EOF reached).")
		os.Exit(1)
	}
	line = scanner.Text()
	return line
}
func main() {
	petSay("Hi, I am momo")
	petSay("My brain is a little squishy")
	fmt.Println()
	for {
		petMenu()
	}
}
