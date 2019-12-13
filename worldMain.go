package main

import {
	"fmt"
	"os"
	"bufio"

type Monster struct {
	Name	string
	Health	int
}

type Player struct {
	Name	string
	Health, Level	int
}

func userOptions(){
	fmt.Println("1. Attack\n2. Defend\n 3. Run\n" )
	choices = 3
	return choices
}

func userDecision(choices int, input string){
	if choices == 3 {
		switch input {
		case 1:
			fmt.Println("You chose to attack")
		case 2:
			fmt.Println("You chose to defend")
		case 3:
			fmt.Println("You chose to run")
		}
	}

	else {
		fmt.Println("Error: Incorrect number of choices")
	}
}

func main() {
	exit = false
	var next string 
	reader := bufio.NewReader(os.Stdin)

	for exit != true {
		next = ""
		fmt.Println("What will you do next?\n")
		numChoices = userOptions()
		next, _ := reader.ReadString('\n')
		if next == 'exit'{
			exit = true
		}

		userDecision(numChoices, next)

		
		
	}
}
