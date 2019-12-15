package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//---------------

//Character Structs
type Monster struct {
	Name           string
	Health, Attack int
	//Monster character for game
}

type Player struct {
	Name                  string
	Health, Attack, Level int
	//Player character
}

//---------------

//Character initialization
func initUser(name string) Player {
	var player = Player{
		Name:   name,
		Health: rand.Intn(20 - 5),
		Attack: rand.Intn(5 - 1),
		Level:  1,
	}
	return player
}

func initMonster() Monster {
	var monster = Monster{
		Name:   "Skeleton",
		Health: rand.Intn(10 - 1),
		Attack: rand.Intn(3 - 0),
	}
	return monster
}

//---------------

// Action Section
func attack(user Player) {
	fmt.Println(user.Name, "attacked for", rand.Intn(user.Attack-0), "damage!")
}

func run(user Player, enemy Monster) {
	if rand.Intn(2-0) == 0 {
		fmt.Println("You got away!")
	} else {
		fmt.Println("You couldn't get away!")
	}

}

//---------------

//Main menu functionality
func userOptions() {
	fmt.Println("1. Attack\n2. Defend\n3. Run\n4. Exit")
}
func userDecision(choices int, input string, user Player, monster Monster) {
	if choices == 3 {
		switch input {
		case "1":
			fmt.Println("You chose to attack the", monster.Name)
			attack(user)

		case "2":
			fmt.Println("You chose to defend")

		case "3":
			fmt.Println("You chose to run")
			run(user, monster)
		default:
			fmt.Println("No options selected")
		}
	} else {
		fmt.Println("Error: Incorrect number of choices")
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

//---------------

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
			userDecision(3, nextChoice, user, monsterObj)
		}

	}
	fmt.Println("Goodbye!")
}
