package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var fileIsCorrupted = false

func saveMemories() error {
	if fileIsCorrupted {
		return fmt.Errorf("refusing to save: memory file was loaded in a corrupted state")
	}
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
		fmt.Printf("%sWarning: Could not read the memory file (%v).Starting fresh.%s\n", ColorRed, err, ColorReset)
		petMemoriesRaw = []Memory{}
		return
	}

	var loadedMemories []Memory
	err = json.Unmarshal(data, &loadedMemories)
	if err != nil {
		fileIsCorrupted = true

		backupName := memoryFile + ".bak"
		_ = os.WriteFile(backupName, data, 0644)
		fmt.Printf("%sCRITICAL ERROR: Memory file '%s' is corrupted! (%v)%s\n", ColorRed, memoryFile, err, ColorReset)
		fmt.Printf("%sBacked up broken file to '%s'.Momo will NOT overwrite your data%s\n", ColorYellow, backupName, ColorReset)
		petMemoriesRaw = []Memory{}
		return
	}
	petMemoriesRaw = loadedMemories
}
