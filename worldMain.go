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

func monsterDecision(monster *monster, user *player) bool {
	if monster.getHealth() > (monster.getHealth() / 2) {
		monsterAttack(monster, user)
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

//Reset main function
func main() {
	var exit bool

	entryMenu := initMenu("start")
	mainMenu := initMenu("main")
	entryMenu.printMenu()
	userName := readUserInput()
	clearScreen()
	user := initUser(userName)
	fmt.Println(user.getName(), "has entered the world...")
	for exit == false {
		fmt.Println("What will  you do next?")
		mainMenu.printMenu()

		nextChoice := readUserInput()
		clearScreen()

		exit = mainAction(nextChoice, user)

	}
}
