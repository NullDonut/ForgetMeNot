package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	color.NoColor = false
	loadMemories()
	petSay("Hi, I am momo")
	petSay("My brain is a little squishy")
	fmt.Println()
	for {
		cleanMemories()
		petMenu()
	}
}
