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
	fmt.Println("1. Attack\n2. Heal\n3. Run\n4. Exit")
}

func menuAction(input string, user *player, monster *monster, menu string) bool {
	//var subExit bool
	switch strings.ToLower(menu) {
	/*case "main":
	switch strings.ToLower(input) {
	case "1":
		fmt.Println("Time to train...")
		var combatMenu = initMenu("combat")
		combatMenu.printMenu()
		for subExit == false {
			subExit == menuAction("combat")
		}
		return false
	case "2":
		fmt.Println("Time to explore...")
		return false
	case "3":
		fmt.Println("Time to rest....")
		return false
	case "4":
		fmt.Println("Goodbye!")
		return true
	default:
		fmt.Println("No options selected, returning to main menu")
		return false
	}*/
	case "combat":
		switch strings.ToLower(input) {
		case "1":
			attack(user, monster)
			if monster.getHealth() <= 0 {
				fmt.Println("You defeated the", monster.getName())
				return true
			}
			return false
		case "2":
			user.heal()
			return false
		case "3":
			return run(user, monster)
		case "4", "exit":
			fmt.Println("Goodbye!")
			return true
		default:
			fmt.Println("No options selected")
			return false
		}
	default:
		fmt.Println("No options selected")
		return false
	}
}

func monsterDecision(monster *monster, user *player) bool {
	if monster.getHealth() > (monster.getHealth() / 2) {
		monsterAttack(monster, user)
		//fmt.Println(user.Name, "health:", user.Health)
		if user.getHealth() <= 0 {
			fmt.Println("You were defeated by the", monster.getName())
			return true
		}
		return false
	} else if monster.getHealth() > (monster.getHealth() / 4) {
		monster.heal()
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

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var exit bool

	entryMenu := initMenu("start")
	combatMenu := initMenu("combat")
	//mainMenu := initMenu("main")
	entryMenu.printMenu()
	userName := readUserInput()
	clearScreen()
	user := initUser(userName)
	monsterObj := initMonster()
	fmt.Println(user.getName(), "has entered the world...")
	for exit == false {
		fmt.Println("What will  you do next?")
		fmt.Println(user.getName(), "health:", user.getHealth(), "\t", monsterObj.getName(), "health:", monsterObj.getHealth())
		combatMenu.printMenu()
		print(">>")
		nextChoice := readUserInput()
		clearScreen()
		exit = menuAction(nextChoice, user, monsterObj, "combat")

		if exit == false {
			exit = monsterDecision(monsterObj, user)
		}
	}
}
