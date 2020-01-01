package main

import (
	"math/rand"
)

//Character struct
type character struct {
	Name           string
	Health, Attack int
}

//Character Interface
type Character interface {
	Name() string
	Health() int
	Attack() int
	getAttack() int
	getHealth() int
	getName() string
	//setHealth(damage int)
}

//Monster struct
type monster struct {
	character
	Type string
}

//Monster Interface
type Monster interface {
	Character
	Type() string
}

//Player struct
type player struct {
	character
	Level int
}

//Player Interface
type Player interface {
	Character
	Level() int
}

func (c *character) getName() string {
	return c.Name
}

func (c *character) getAttack() int {
	return c.Attack
}

func (c *character) getHealth() int {
	return c.Health
}

func (c *character) setHealth(damage int) {
	c.Health = c.getHealth() - damage
}

func initUser(name string) *player {
	player := &player{
		character: character{Name: name, Health: rand.Intn(20 - 5), Attack: rand.Intn(5 - 1)},
		Level:     1,
	}
	return player
}

func initMonster() *monster {
	monster := &monster{
		character: character{Health: rand.Intn(10 - 1), Attack: rand.Intn(3 - 0)},
		Type:      "Skeleton",
	}
	return monster
}
