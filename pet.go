package main

import (
	"fmt"
	"math"
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
	petSay("I'll try to remember that!")
	var m = Memory{Fact: Fact, MemoryStrength: 1.0, TimeLearned: time.Now()}
	petMemoriesRaw = append(petMemoriesRaw, m)
	saveMemories()
}
