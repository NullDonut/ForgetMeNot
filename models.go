package main

import "time"

type Memory struct {
	Fact        string    `json:"fact"`
	TimeLearned time.Time `json:"time_learned"`
	Stability   float64   `json:"stability"`
	ReviewCount int       `json:"review_count"`
}

const memoryFile = "momo_memories.json"
const petName string = "Momo"
