package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Monster struct {
	Name	string
	Health	int
}

type Player struct {
	Name	string
	Health, Level	int
}

func userOptions() {
	fmt.Println("1. Attack\n2. Defend\n3. Run\n4. Exit")
}

func userDecision(choices int, input string){
	if choices == 3 {
		switch input { 
		case "1":
			fmt.Println("You chose to attack")
		case "2":
			fmt.Println("You chose to defend")
		case "3":
			fmt.Println("You chose to run")
		default:
			fmt.Println("No options selected")
		}
	}	else {
		fmt.Println("Error: Incorrect number of choices")
	}
}

func readUserInput() string{
	reader := bufio.NewReader(os.Stdin)
	next, err := reader.ReadString('\n')
		if err != nil {	
			fmt.Fprintln(os.Stderr, err)
		}
	return strings.TrimRight(next, "\r\n")
}

func main() {
	var exit bool
	for exit == false {
		fmt.Println("What will you do next?")
		userOptions()
		print(">>")
		nextChoice := readUserInput()
		if nextChoice == "exit" || nextChoice == "4"{
			exit = true
		} else{
			userDecision(3, nextChoice)
		}
		
	}
	fmt.Println("Goodbye!")
}
