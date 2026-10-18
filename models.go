package main

import "time"

type Memory struct {
	Fact           string    `json:"fact"`
	MemoryStrength float64   `json:"memory_strength"`
	TimeLearned    time.Time `json:"time_learned"`
}

const memoryFile = "momo_memories.json"
const petName string = "MoMO"
