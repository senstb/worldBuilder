package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Main menu functionality
func userOptions() {
	fmt.Println("1. Attack\n2. Defend\n3. Run\n4. Exit")
}
func userDecision(input string, user Player, monster Monster) {
	switch input {
	case "1":
		fmt.Println("You chose to attack the", monster.Name)
		attack(user, monster)
		fmt.Println("Monster Health:", monster.Health)
		if monster.Health <= 0 {
			fmt.Println("You defeated the", monster.Name)
		}

	case "2":
		fmt.Println("You chose to defend")

	case "3":
		fmt.Println("You chose to run")
		run(user, monster)
	default:
		fmt.Println("No options selected")
	}
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	next, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return strings.TrimRight(next, "\r\n")
}

func clearScreen() {
	exec.Command("cmd", "/c", "cls")
}

func main() {
	var exit bool
	fmt.Println("Who is the next adventurer?")
	userName := readUserInput()
	user := initUser(userName)
	monsterObj := initMonster()
	fmt.Println(user.Name, "has entered the world...")
	for exit == false {
		fmt.Println("What will  you do next?")
		userOptions()
		print(">>")
		nextChoice := readUserInput()
		if nextChoice == "exit" || nextChoice == "4" {
			exit = true
		} else {
			userDecision(nextChoice, user, monsterObj)
		}
	}
	fmt.Println("Goodbye!")
}
