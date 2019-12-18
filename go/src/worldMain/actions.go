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

func run(user *Player, monster *Monster) bool {
	if rand.Intn(2-0) == 0 {
		fmt.Println(user.Name, "got away from the", monster.Name)
		return true
	}
	fmt.Println(user.Name, "couldn't get away from the", monster.Name)
	return false
}
func monsterAttack(monster *Monster, user *Player) {
	attackValue := rand.Intn(monster.Attack - 0)
	fmt.Println(monster.Name, "attacked for", attackValue, "damage!")
	monster.Health = user.Health - attackValue
}

func monsterDefend(monster *Monster, user *Player) {
	defendValue := rand.Intn(monster.Attack - 1)
	fmt.Println("You defend for", defendValue, "damage!")
}
func monsterRun(monster *Monster, user *Player) bool {
	if rand.Intn(2-0) == 0 {
		fmt.Println(monster.Name, "got away from", user.Name)
		return true
	}
	fmt.Println(monster.Name, "couldn't get away from", user.Name)
	return false
}
