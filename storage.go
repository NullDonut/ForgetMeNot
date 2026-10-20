package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveMemories() error {
	data, err := json.MarshalIndent(petMemoriesRaw, "", "	")
	if err != nil {
		return fmt.Errorf("failed to marshal memories: %w", err)
	}
	tmpFileName := memoryFile + ".tmp"
	err = os.WriteFile(tmpFileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write temp memory file:%w", err)
	}

	err = os.Rename(tmpFileName, memoryFile)
	if err != nil {
		os.Remove(tmpFileName)
		return fmt.Errorf("failed to replace memory file: %w", err)
	}
	return nil
}
func loadMemories() {
	if _, err := os.Stat(memoryFile); os.IsNotExist(err) {
		petMemoriesRaw = []Memory{}
		return
	}
	data, err := os.ReadFile(memoryFile)
	if err != nil {
		fmt.Printf("%sWarning: Could not read the memory file (%v).Starting fres.%s\n", ColorRed, err, ColorReset)
		petMemoriesRaw = []Memory{}
		return
	}

	var loadedMemories []Memory
	err = json.Unmarshal(data, &loadedMemories)
	if err != nil {
		fmt.Printf("%CRITICAL ERROR: Memory file '%s' is corrupted! (%v)%s\n", ColorRed, memoryFile, err, ColorReset)
		fmt.Printf("%sTo protect your data,Momo will not overwrite this file. Please inspect or back up '%s'.%s\n", ColorYellow, memoryFile, ColorReset)
		petMemoriesRaw = []Memory{}
		return
	}
	petMemoriesRaw = loadedMemories
}
