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

func userDecision(input string, user *player, monster *monster) bool {
	switch strings.ToLower(input) {
	case "1":
		attack(user, monster)
		//fmt.Println("Monster Health:", monster.Health)
		if monster.getHealth() <= 0 {
			fmt.Println("You defeated the", monster.getName())
			return true
		}
		return false
	case "2":
		heal(user, monster)
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
		monsterHeal(monster, user)
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
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var exit bool
	fmt.Println("Who is the next adventurer?")
	userName := readUserInput()
	clearScreen()
	user := initUser(userName)
	monsterObj := initMonster()
	fmt.Println(user.getName(), "has entered the world...")
	for exit == false {
		//		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Println(user.getName(), "health:", user.getHealth(), "\t", monsterObj.getName(), "health:", monsterObj.getHealth())
		fmt.Println("What will  you do next?")
		userOptions()
		print(">>")
		nextChoice := readUserInput()
		clearScreen()
		exit = userDecision(nextChoice, user, monsterObj)
		if exit == false {
			exit = monsterDecision(monsterObj, user)
		}
	}
}
