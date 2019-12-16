package main

import (
	"fmt"
	"math/rand"
)

func attack(user Player, monster Monster) {
	attackValue := rand.Intn(user.Attack - 1)
	fmt.Println(user.Name, "attacked for", attackValue, "damage!")
	monster.Health = monster.Health - attackValue
}

func run(user Player, enemy Monster) {
	if rand.Intn(2-0) == 0 {
		fmt.Println("You got away from the", enemy)
	} else {
		fmt.Println("You couldn't get away from the", enemy)
	}

}
