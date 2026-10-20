package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	ColorReset  = "\033[m"
	ColorCyan   = "\033[036m"
	ColorGreen  = "\033[032m"
	ColorYellow = "\033[033m"
	ColorPurple = "\033[035m"
	ColorRed    = "\033[031m"
)

func petMashup() {
	if len(petMemoriesRaw) < 2 {
		return
	}

	idx1 := rand.Intn(len(petMemoriesRaw))
	idx2 := rand.Intn(len(petMemoriesRaw))

	for idx1 == idx2 {
		idx2 = rand.Intn(len(petMemoriesRaw))
	}

	fact1 := petMemoriesRaw[idx1].Fact
	fact2 := petMemoriesRaw[idx2].Fact

	words1 := strings.Fields(fact1)
	words2 := strings.Fields(fact2)

	if len(words1) == 0 || len(words2) == 0 {
		return
	}
	mid1 := len(words1) / 2
	if mid1 == 0 {
		mid1 = 1
	}
	mid2 := len(words2) / 2

	firstHalf := strings.Join(words1[:mid1], " ")
	secondHalf := strings.Join(words2[mid2:], " ")

	fmt.Println()
	petSay(fmt.Sprintf("%s*scratches head* ...Wait! Is it true that %s %s...? My brain mixed up those up! %s", ColorPurple, firstHalf, secondHalf, ColorReset))
	fmt.Println()
}
func petDream() {
	if len(petMemoriesRaw) == 0 {
		return
	}

	idx := rand.Intn(len(petMemoriesRaw))
	mem := petMemoriesRaw[idx]

	fmt.Println()
	petSay(fmt.Sprintf("%s*yawns and rubs eyes* ...zzZ... Oh! I was just dreaming about.. \"%s\" ...zzZ%s", ColorPurple, mem.Fact, ColorReset))
}
func getPetLevel() (int, float64) {
	totalStrength := 0.0
	activeCount := 0

	for _, m := range petMemoriesRaw {
		strength := MemoryStrength(m)
		if strength > 0.05 {
			totalStrength += strength
			activeCount++
		}
	}
	level := 1 + int(totalStrength/3.0)
	return level, totalStrength
}
func GetStrengthBar(strength float64) string {
	totalBlocks := 10
	filled := int(strength * float64(totalBlocks))
	if filled > totalBlocks {
		filled = totalBlocks
	}
	if filled < 0 {
		filled = 0
	}
	bar := ""
	for i := 0; i < totalBlocks; i++ {
		if i < filled {
			bar += "▇"

		} else {
			bar += "-"
		}
	}
	return fmt.Sprintf("[%s] %.0f%%", bar, strength*100)
}
func MemoryStrength(value Memory) float64 {
	k := 0.1
	timeElapsed := time.Since(value.TimeLearned).Hours()
	currentStrength := (value.MemoryStrength) * math.Pow(math.E, -1*timeElapsed*k)
	return currentStrength
}
func cleanMemories() {
	var aliveMemories []Memory
	for _, m := range petMemoriesRaw {
		if MemoryStrength(m) > 0.05 {
			aliveMemories = append(aliveMemories, m)
		}
	}
	petMemoriesRaw = aliveMemories
	saveMemories()
}

func memoryReinforce(mem *Memory) {
	mem.MemoryStrength = 1
	mem.TimeLearned = time.Now()
}

func petSay(msg string) {
	fmt.Printf("%s[Momo]:%s %s\n", ColorCyan, ColorReset, msg)
}
func petUserInputFeedback(Fact string) {
	trimmed := strings.TrimSpace(Fact)
	if trimmed == "" {
		petSay(fmt.Sprintf("%sYou didn't teach me anything! Give me actual words to remember.%s", ColorRed, ColorReset))
		return
	}
	petSay(fmt.Sprintf("%sI'll try to remember that!%s", ColorGreen, ColorReset))
	var m = Memory{Fact: Fact, MemoryStrength: 1.0, TimeLearned: time.Now()}
	petMemoriesRaw = append(petMemoriesRaw, m)
	saveMemories()
}
