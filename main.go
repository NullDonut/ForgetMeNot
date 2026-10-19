package main

import (
	"fmt"
)

func main() {
	loadMemories()
	petDream()
	petSay("Hi, I am momo")
	petSay("My brain is a little squishy")
	fmt.Println()
	for {
		cleanMemories()
		petMenu()
	}
}
