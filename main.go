package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Memory struct {
	fact           string
	memoryStrength float64
	timeLearned    time.Time
}

const petName string = "Momo"

var scanner = bufio.NewScanner(os.Stdin)
var petMemoriesRaw []Memory

func petSay(msg string) {
	fmt.Printf("%s: %s\n", petName, msg)
}
func memoryReinforce(mem *Memory) {
	mem.memoryStrength = 1
	mem.timeLearned = time.Now()
}
func petUserInputFeedback(fact string) {
	petSay("I'll try to remember that!")
	// petSay(fact)
	var m = Memory{fact: fact, memoryStrength: 1.0, timeLearned: time.Now()}
	petMemoriesRaw = append(petMemoriesRaw, m)
	// fmt.Printf("You have this much memories: %d\n", len(petMemoriesRaw))
}
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
		fact := inputUser()
		petUserInputFeedback(fact)
	case "2":
		if (len(petMemoriesRaw)) == 0 {
			petSay("I don't remember anything yet :( ,Teach me something")
			return
		}
		petSay("Let me see what I remember!")
		for _, value := range petMemoriesRaw {

			petSay(strconv.FormatFloat(memoryStrength(value), 'f', 6, 64))
			petSay(value.timeLearned.Format("2006-01-02 15:04:05"))
		}
	case "3":
		if (len(petMemoriesRaw)) == 0 {
			petSay("I don't remember anything yet :( ,Teach me something")
			return
		}
		petSay("My memories i still remember a bit")
		for i, value := range petMemoriesRaw {
			fmt.Printf("%d: %s\n", i+1, value.fact)
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
func memoryStrength(value Memory) float64 {
	k := 0.1
	timeElapsed := time.Since(value.timeLearned).Hours()
	currentStrength := (value.memoryStrength) * math.Pow(math.E, -1*k*timeElapsed)
	return currentStrength
}
func main() {
	petSay("Hi, I am momo")
	petSay("My brain is a little squishy")
	fmt.Println()
	for {
		petMenu()
	}
}
