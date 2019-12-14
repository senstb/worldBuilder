package main

import (
	"fmt"
	"os"
	"bufio"
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
	fmt.Println("1. Attack\n2. Defend\n3. Run" )
	//var choices = 3
}


func userDecision(choices int, input string){
//	fmt.Println("This is the userDecision Function input: " + input)

	if choices == 3 {
		switch input { 
		case "a":
			fmt.Println("You chose to attack")
		case "b":
			fmt.Println("You chose to defend")
		case "":
			fmt.Println("You chose to run")
		default:
			fmt.Println("No options selected")
		}
	}	else {
		fmt.Println("Error: Incorrect number of choices")
	}

}


func main() {
	reader := bufio.NewReader(os.Stdin)
	var next string
	var err error
	for true {
		fmt.Println("What will you do next?")
		userOptions()
		next, err = reader.ReadString('\n')
		if err != nil {	
			fmt.Fprintln(os.Stderr, err)
		}
		userDecision(3, next)
	}
	fmt.Println("Goodbye!")
}
