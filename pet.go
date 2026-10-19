package main

import (
	"github.com/fatih/color"
	"math"
	"time"
)

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
	color.Cyan("[Momo]: %s\n", msg)
}
func petUserInputFeedback(Fact string) {
	petSay("I'll try to remember that!")
	var m = Memory{Fact: Fact, MemoryStrength: 1.0, TimeLearned: time.Now()}
	petMemoriesRaw = append(petMemoriesRaw, m)
	saveMemories()
}
