package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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
		fmt.Printf("%sCould not read: '%s' : %v%s\n", ColorRed, memoryFile, err, ColorReset)
		os.Exit(1)
	}

	var loadedMemories []Memory
	err = json.Unmarshal(data, &loadedMemories)
	if err != nil {
		backup := fmt.Sprintf("%s.%d.bak", memoryFile, time.Now().Unix())
		_ = os.WriteFile(backup, data, 0644)
		fmt.Printf("%sMemory file is corrupted (%v). Backed up to %s.Fix or delete it,then restart.%s\n", ColorRed, err, backup, ColorReset)
		os.Exit(1)
	}
	petMemoriesRaw = loadedMemories
}
