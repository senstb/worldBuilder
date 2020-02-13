package main

import (
	"fmt"
	"strings"
)

type menu struct {
	Context, Options string
	OptionCount      int
}

func (m *menu) getContext() string {
	return m.Context
}

func (m *menu) getOptionCount() int {
	return m.OptionCount
}

func (m *menu) getOptions() string {
	return m.Options
}

func (m *menu) printMenu() {
	fmt.Println(m.Options)
	print(">>")
}

func combatAction(input string, user *player, monster *monster, menu string) bool {
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
}

func mainAction(input string, user *player) bool {
	origHealth := user.getHealth()
	switch strings.ToLower(input) {
	case "1":
		monster := initMonster()
		var subExit bool
		//origHealth := user.getHealth()
		for subExit == false {
			var combatMenu = initMenu("combat")
			fmt.Println(user.getName(), "health:", user.getHealth(), "\t", monster.getName(), "health:", monster.getHealth())
			combatMenu.printMenu()
			nextChoice := readUserInput()
			clearScreen()
			subExit = combatAction(nextChoice, user, monster, "combat")
			if subExit == false {
				subExit = monsterDecision(monster, user)
			}
		}
		user.setHealth(origHealth)
		return false
	case "2":
		fmt.Println("Time to explore...")
		return false
	case "3":
		fmt.Println("Time to rest....")
		user.setHealth(origHealth)
		fmt.Println("Health restored")
		return false
	case "4":
		fmt.Println("Goodbye!")
		return true
	default:
		fmt.Println("No options selected, returning to main menu")
		return false
	}
}

func initMenu(situation string) menu {
	switch strings.ToLower(situation) {
	case "combat":
		menuOutput := menu{
			Context:     situation,
			Options:     "1. Attack\n2. Heal\n3. Run",
			OptionCount: 3,
		}
		return menuOutput
	case "start":
		menuOutput := menu{
			Context:     situation,
			Options:     "Who is the next adventurer?",
			OptionCount: 1,
		}
		return menuOutput
	default:
		menuOutput := menu{
			Context:     "main",
			Options:     "1. Train\n2. Explore\n3. Rest\n4. Exit",
			OptionCount: 4,
		}
		return menuOutput
	}

}
