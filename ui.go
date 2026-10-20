package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var scanner = bufio.NewScanner(os.Stdin)
var petMemoriesRaw []Memory

func normalize(s string) string {
	return strings.TrimFunc(strings.ToLower(strings.TrimSpace(s)),
		func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
}
func runQuiz(mem *Memory) {
	petSay("Quiz time! Can you complete this memory?")
	fmt.Println("------------------------------------------------------")

	words := strings.Fields(mem.Fact)
	answer := mem.Fact
	var candidates []int
	for i, w := range words {
		if normalize(w) != "" {
			candidates = append(candidates, i)
		}
	}
	if len(words) > 1 && len(candidates) > 0 {
		i := candidates[rand.Intn(len(candidates))]
		answer = words[i]
		words[i] = "____"
		fmt.Println("Prompt:", strings.Join(words, " "))
	} else {
		r := []rune(answer)
		fmt.Printf("Starts with '%c' , %d letters\n", r[0], len(r))
	}
	userAnswer := normalize(inputUser())
	if userAnswer != "" && userAnswer == normalize(answer) {
		mult := memoryReinforce(mem)
		if mult == 1.0 {
			petSay("Still fresh in my head! Ask me again later.")
		} else {
			petSay(fmt.Sprintf("%sCorrect! Stability *%.2f %s", ColorGreen, mult, ColorReset))
		}
	} else {
		memoryLapse(mem)
		petSay(fmt.Sprintf("%sNot quite! The answer was '%s'. (Stability dropped)%s", ColorRed, answer, ColorReset))
	}
	if err := saveMemories(); err != nil {
		fmt.Printf("%sFailed to save: %v%s\n", ColorRed, err, ColorReset)
	}
	fmt.Println("------------------------------------------------------")
}

func petMenu() {
	switch r := rand.Float64(); {
	case r < 0.10:
		petMashup()
	case r < 0.15:
		petDreamGarbled()
	}
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
		mem := getWeakestMemory()
		if mem == nil {
			petSay("I don't remember anything yet ;(Teach me something first!")
			return
		}
		runQuiz(mem)
	case "4":
		petSay("Byeee 1 Come back before I forget you!")
		os.Exit(0)
	default:
		petSay(fmt.Sprintf("%sRetry plss....%s", ColorRed, ColorReset))
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
