package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Main menu functionality
func userOptions() {
	fmt.Println("1. Attack\n2. Defend\n3. Run\n4. Exit")
}

func userDecision(input string, user *Player, monster *Monster) bool {
	switch strings.ToLower(input) {
	case "1":
		fmt.Println("You chose to attack the", monster.Name)
		attack(user, monster)
		fmt.Println("Monster Health:", monster.Health)
		if monster.Health <= 0 {
			fmt.Println("You defeated the", monster.Name)
			return true
		}
		return false
	case "2":
		fmt.Println("You chose to defend")
		defend(user, monster)
		return false
	case "3":
		fmt.Println("You chose to run")
		return run(user, monster)
	case "4", "exit":
		fmt.Println("Goodbye!")
		return true
	default:
		fmt.Println("No options selected")
		return false
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
		exit = userDecision(nextChoice, &user, &monsterObj)
	}
}
