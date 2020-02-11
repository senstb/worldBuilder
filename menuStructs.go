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
