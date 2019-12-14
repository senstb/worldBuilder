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
	fmt.Println("1. Attack\n2. Defend\n3. Run" )
	//var choices = 3
}


func userDecision(choices int, input string){
//	fmt.Println("This is the userDecision Function input: " + input)

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


func main() {
	reader := bufio.NewReader(os.Stdin)
	var next string
	var err error
	var exit bool
	for exit == false {
		fmt.Println("What will you do next?")
		userOptions()
		print(">>")
		next, err = reader.ReadString('\n')
		if err != nil {	
			fmt.Fprintln(os.Stderr, err)
		}
		if strings.TrimRight(next, "\r\n") == "exit"{
			exit = true
		} else{
			userDecision(3, strings.TrimRight(next, "\r\n"))
		}
		

	}
	fmt.Println("Goodbye!")
}
