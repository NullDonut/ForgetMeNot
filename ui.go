package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var petMemoriesRaw []Memory

func petMenu() {
	fmt.Println()
	petSay("What should we do? ")
	fmt.Println("1. Teach me something")
	fmt.Println("2. What do you remember?")
	fmt.Println("3. Quiz me")
	fmt.Println("4. Quit")
	var choice string
	fmt.Println()
	fmt.Printf("Choice[1/2/3/4]: ")
	if scanner.Scan() {
		choice = scanner.Text()
	}
	fmt.Println()
	switch choice {
	case "1":
		petSay("Learning time....")
		Fact := inputUser()
		petUserInputFeedback(Fact)
	case "2":
		if (len(petMemoriesRaw)) == 0 {
			petSay("I don't remember anything yet :( ,Teach me something")
			return
		}
		petSay("Let me see what I remember!")
		activeMemoryCount := 0
		for _, value := range petMemoriesRaw {
			strength := MemoryStrength(value)
			if strength > 0.05 {

				petSay(strconv.FormatFloat(MemoryStrength(value), 'f', 6, 64))
				petSay(value.TimeLearned.Format("2006-01-02 15:04:05"))
				activeMemoryCount++
			}
		}
		if activeMemoryCount == 0 {
			petSay("Oh no... I've forgotten all my memories completely! You need to teach me again.")
		}
	case "3":
		if (len(petMemoriesRaw)) == 0 {
			petSay("I don't remember anything yet :( ,Teach me something")
			return
		}
		petSay("My memories i still remember a bit")
		for i, value := range petMemoriesRaw {
			fmt.Printf("%d: %s\n", i+1, value.Fact)
		}
		fmt.Printf("Chose a memory to reinforce: ")
		choiceIdx := inputUser()
		choiceIdxInt, err := strconv.Atoi(choiceIdx)
		if err != nil {
			fmt.Println("Uhh you chose a wrong memory.... Uhhhh.... I am failing.....")
			return
		}
		if !(0 < choiceIdxInt && choiceIdxInt <= len(petMemoriesRaw)) {
			fmt.Println("Uhh i don't remember it... :( Try again")
			return
		}
		memoryReinforce(&petMemoriesRaw[choiceIdxInt-1])
		saveMemories()
	case "4":
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
