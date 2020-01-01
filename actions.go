package main

import (
	"fmt"
	"math/rand"
)

func attack(user *player, monster *monster) {
	attackValue := rand.Intn(user.getAttack() - 1)
	fmt.Println(user.getName(), "attacked for", attackValue, "damage!")
	monster.setHealth(attackValue)
}

func heal(user *player, monster *monster) {
	healValue := rand.Intn(user.getAttack()-1) * -1
	fmt.Println("You healed for", healValue, "damage!")
	user.setHealth(healValue)
}

func run(user *player, monster *monster) bool {
	if rand.Intn(2-0) == 0 {
		fmt.Println(user.getName(), "got away from the", monster.getName())
		return true
	}
	fmt.Println(user.getName(), "couldn't get away from the", monster.getName())
	return false
}

func monsterAttack(monster *monster, user *player) {
	attackValue := rand.Intn(user.getAttack() - 1)
	fmt.Println(monster.getName(), "attacked for", attackValue, "damage!")
	user.setHealth(attackValue)
}

func monsterHeal(monster *monster, user *player) {
	healValue := rand.Intn(monster.getAttack()-1) * -1
	fmt.Println(monster.getName(), "healed for", healValue, "damage!")
	monster.setHealth(healValue)
}

func monsterRun(monster *monster, user *player) bool {
	if rand.Intn(2-0) == 0 {
		fmt.Println(monster.getName(), "got away from", user.getName())
		return true
	}
	fmt.Println(monster.getName(), "couldn't get away from", user.getName())
	return false
}
