package main

import (
	"fmt"
	"math/rand"
)

func attack(user *Player, monster *Monster) {
	attackValue := rand.Intn(user.Attack - 1)
	fmt.Println(user.Name, "attacked for", attackValue, "damage!")
	monster.Health = monster.Health - attackValue
}

func defend(user *Player, monster *Monster) {
	defendValue := rand.Intn(monster.Attack - 1)
	fmt.Println("You defend for", defendValue, "damage!")
}

func run(user *Player, enemy *Monster) bool {
	if rand.Intn(2-0) == 0 {
		fmt.Println("You got away from the", enemy.Name)
		return true
	}
	fmt.Println("You couldn't get away from the", enemy.Name)
	return false
}
