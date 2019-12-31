package main

import (
	"math/rand"
)

//Character struct
type Character struct {
	Name           string
	Health, Attack int
}

//Monster struct
type Monster struct {
	Character
	Type string
}

//Player struct
type Player struct {
	Character
	Level int
}

func (c *Character) getName() string {
	return c.Name
}

func (c *Character) getAttack() int {
	return c.Attack
}

func (c *Character) getHealth() int {
	return c.Health
}

func initUser(name string) Player {
	var player = Player{
		Character: Character{Name: name, Health: rand.Intn(20 - 5), Attack: rand.Intn(5 - 1)},
		Level:     1,
	}
	return player
}

func initMonster() Monster {
	var monster = Monster{
		Character: Character{Health: rand.Intn(10 - 1), Attack: rand.Intn(3 - 0)},
		Type:      "Skeleton",
	}
	return monster
}
