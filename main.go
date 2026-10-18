package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Memory struct {
	Fact           string    `json:"fact"`
	MemoryStrength float64   `json:"memory_strength"`
	TImeLearned    time.Time `json:"time_learned"`
}

const memoryFile = "momo_memories.json"

const petName string = "Momo"

var scanner = bufio.NewScanner(os.Stdin)
var petMemoriesRaw []Memory

func cleanMemories() {
	var aliveMemories []Memory
	for _, m := range petMemoriesRaw {
		if MemoryStrength(m) > 0.05 {
			aliveMemories = append(aliveMemories, m)
		}
		petMemoriesRaw = aliveMemories
		saveMemories()
	}
}
func saveMemories() {
	data, err := json.MarshalIndent(petMemoriesRaw, "", "	")
	if err != nil {
		return
	}
	_ = os.WriteFile(memoryFile, data, 0644)
}
func loadMemories() {
	data, err := os.ReadFile(memoryFile)
	if err != nil {
		return
	}
	_ = json.Unmarshal(data, &petMemoriesRaw)
}

func petSay(msg string) {
	fmt.Printf("%s: %s\n", petName, msg)
}
func memoryReinforce(mem *Memory) {
	mem.MemoryStrength = 1
	mem.TImeLearned = time.Now()
}
func petUserInputFeedback(Fact string) {
	petSay("I'll try to remember that!")
	// petSay(Fact)
	var m = Memory{Fact: Fact, MemoryStrength: 1.0, TImeLearned: time.Now()}
	petMemoriesRaw = append(petMemoriesRaw, m)
	saveMemories()
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
				petSay(value.TImeLearned.Format("2006-01-02 15:04:05"))
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
func MemoryStrength(value Memory) float64 {
	k := 0.1
	timeElapsed := time.Since(value.TImeLearned).Hours()
	currentStrength := (value.MemoryStrength) * math.Pow(math.E, -1*k*timeElapsed)
	return currentStrength
}
func main() {
	loadMemories()
	petSay("Hi, I am momo")
	petSay("My brain is a little squishy")
	fmt.Println()
	for {
		cleanMemories()
		petMenu()
	}
}
