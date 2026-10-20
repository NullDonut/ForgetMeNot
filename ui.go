package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	if !scanner.Scan() {
		fmt.Println("\nGoodbye!")
		os.Exit(0)
	}
	choice = strings.TrimSpace(scanner.Text())
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
		idx := rand.Intn(len(petMemoriesRaw))
		mem := &petMemoriesRaw[idx]
		petSay("Qioz time! Let's see if we can recall this memory: ")
		fmt.Println("------------------------------------------------------")
		words := strings.Fields(mem.Fact)
		if len(words) > 1 {
			hiddenWord := words[len(words)-1]
			words[len(words)-1] = "____"
			fmt.Printf("Hint/Prompt: %s\n", strings.Join(words, " "))

			fmt.Printf("%sType the missing word or full fact >%s", ColorYellow, ColorReset)
			userAnswer := inputUser()

			if strings.EqualFold(strings.TrimSpace(userAnswer), hiddenWord) || strings.EqualFold(strings.TrimSpace(userAnswer), mem.Fact) {
				memoryReinforce(mem)
				if err := saveMemories(); err != nil {
					fmt.Printf("%sFailed to save memories: %v%s\n", ColorRed, err, ColorReset)
				}
				petSay(fmt.Sprintf("%s Correct! Memory succesfully reinforced! (Stability doubled!)%s", ColorGreen, ColorReset))
			} else {
				petSay(fmt.Sprintf("%s Ah,not quite! The correct answer was '%s'.Let's study harder next time.%s", ColorRed, hiddenWord, ColorReset))
			}
		} else {
			fmt.Printf("Fact: %s \n", mem.Fact)
			fmt.Printf("%sType it back to confirm >%s ", ColorYellow, ColorReset)
			userAnswer := inputUser()
			if strings.EqualFold(strings.TrimSpace(userAnswer), mem.Fact) {
				memoryReinforce(mem)
				if err := saveMemories(); err != nil {
					fmt.Printf("%sFailed to save memories: %v%s\n", ColorRed, err, ColorReset)
				}
				petSay(fmt.Sprintf("%s Nailed it!%s", ColorGreen, ColorReset))
			} else {
				petSay(fmt.Sprintf("%s Oops,that wasn't quite right.%s", ColorRed, ColorReset))
			}
		}
		fmt.Println("------------------------------------------------------")
	case "4":
		petSay("Byeee 1 Come back before I forget you!")
		os.Exit(0)
	default:
		petSay(fmt.Sprintf("%sI don't understand your choie ;( Retry plss....%s", ColorRed, ColorReset))
		petMashup()
	}
}
func inputUser() string {
	fmt.Printf("> ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading line:", err)
		}
		fmt.Println("\nNo Input Provided (EOF reached).")
		os.Exit(0)
	}
	return strings.TrimSpace(scanner.Text())
}
