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
	level, totalStrength := getPetLevel()
	fmt.Println("\n====================================================")
	fmt.Printf("	%s[Momo - Level %d]%s | Brain Power:%.3f\n", ColorPurple, level, ColorReset, totalStrength)
	petSay("What should we do? ")
	fmt.Println("------------------------------------------------------")
	fmt.Printf("	%s1.%s Teach me something\n", ColorCyan, ColorReset)
	fmt.Printf("	%s2.%s What do you remember?\n", ColorCyan, ColorReset)
	fmt.Printf("	%s3.%s Quiz me / Reinforce memory\n", ColorCyan, ColorReset)
	fmt.Printf("	%s4.%s Quit\n", ColorCyan, ColorReset)
	fmt.Println("====================================================")

	var choice string
	fmt.Printf("%sChoice [1-4] >%s ", ColorYellow, ColorReset)
	if scanner.Scan() {
		choice = scanner.Text()
	}
	fmt.Println()
	switch choice {
	case "1":
		petSay("Learning time.... Type what you want me to know:")
		Fact := inputUser()
		petUserInputFeedback(Fact)
	case "2":
		if (len(petMemoriesRaw)) == 0 {
			petSay(fmt.Sprintf("%sI don't remember anything yet :( ,Teach me something!%s", ColorRed, ColorReset))
			return
		}

		fmt.Println("------------------------------------------------------")
		petSay("Let me see what I remember!")
		fmt.Println("------------------------------------------------------")
		activeMemoryCount := 0
		for i, value := range petMemoriesRaw {
			strength := MemoryStrength(value)
			if strength > 0.05 {
				bar := GetStrengthBar(strength)
				fmt.Printf("%s[%d]%s %s\n", ColorYellow, i+1, ColorReset, value.Fact)
				fmt.Printf("	Strength: %s%s%s\n", ColorGreen, bar, ColorReset)
				fmt.Printf("	Learned: %s\n", value.TimeLearned.Format("Jan 2,2006 at 15:04"))
				//fmt.Println("------------------------------------------------------")

				// fmt.Println(GetStrengthBar(MemoryStrength(value)))
				// petSay(value.TimeLearned.Format("2006-01-02 15:04:05"))
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
		petSay("My memories... I  still remember a bit of these:")
		fmt.Println("------------------------------------------------------")
		for i, value := range petMemoriesRaw {
			strength := MemoryStrength(value)
			if strength > 0.05 {
				fmt.Printf("%s[%d]%s %s (Strength: %.0f%%)\n", ColorYellow, i+1, ColorReset, value.Fact, strength*100)
			}
		}
		fmt.Println("------------------------------------------------------")
		fmt.Printf("%sChose a memory to reinforce >%s ", ColorYellow, ColorReset)
		choiceIdx := inputUser()
		choiceIdxInt, err := strconv.Atoi(choiceIdx)
		if err != nil {
			fmt.Println("Uhh you chose a wrong index..... My brain hurts...")
			return
		}
		if !(0 < choiceIdxInt && choiceIdxInt <= len(petMemoriesRaw)) {
			fmt.Println("Uhh i don't remember a memory with that number... :( Try again")
			return
		}
		memoryReinforce(&petMemoriesRaw[choiceIdxInt-1])
		saveMemories()
		petSay(fmt.Sprintf("%s Yay! That memmory is refreshed back to 100%%! %s", ColorGreen, ColorReset))
	case "4":
		petSay("Byeee 1 Come back before I forget you!")
		os.Exit(0)
	default:
		petSay(fmt.Sprintf("%sI don't understand your choie ;( Retry plss....%s", ColorRed, ColorReset))
		petMashup()
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
