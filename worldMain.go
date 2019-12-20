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
		// Monster health not staying static
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

//Needs more work to finalize monsterDecision making and actions
func monsterDecision(monster *Monster, user *Player) bool {
	if monster.Health > (monster.Health / 2) {
		monsterAttack(monster, user)
		fmt.Println(user.Name, "health:", user.Health)
		if user.Health <= 0 {
			fmt.Println("You were defeated by the", monster.Name)
			return true
		}
		return false
	} else if monster.Health > (monster.Health / 4) {
		monsterDefend(monster, user)
		return false
	} else {
		return monsterRun(monster, user)
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
		exit = monsterDecision(&monsterObj, &user)
	}
}
