package main

import (
	"encoding/json"
	"os"
)

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
