package main

import (
	"math/rand"
)

//Monster struct
type Monster struct {
	Name           string
	Health, Attack int
	//Monster character for game
}

//Player struct
type Player struct {
	Name                  string
	Health, Attack, Level int
}

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
