package lib

import (
	"math/rand"
)


type Monster struct {
	Name	string
	Health, Attack	int
}

type Player struct {
	Name	string
	Health, Attack, Level	int
}


func initUser(name string) Player{
	var player = Player{
		Name:	name,
		Health:	rand.Intn(20-5),
		Attack:	rand.Intn(5-0),
		Level:	1,
	}
	return player
}