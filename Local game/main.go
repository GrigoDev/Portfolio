package main

import (
	"fmt"
)

type Hero struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func (h *Hero) Attackayu() {
	fmt.Printf("%s нанёс урон: %d\n", h.Name, h.Attack)
}

func (h *Hero) Defenseyu(damage int) {
	blocked := h.Defense
	if blocked >= damage {
		blocked = damage
	}
	h.Health -= (damage - blocked)
	if h.Health <= 0 {
		h.Health = 0
	}
	fmt.Printf("%s заблокировал урон: %d. Осталось хп: %d.\n", h.Name, blocked, h.Health)
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func (e *Enemy) Attackayu() {
	fmt.Printf("%s нанёс урон: %d\n", e.Name, e.Attack)
}

func (e *Enemy) Defenseyu(damage int) {
	blocked := e.Defense
	if blocked >= damage {
		blocked = damage
	}
	e.Health -= (damage - blocked)
	if e.Health <= 0 {
		e.Health = 0
	}
	fmt.Printf("%s заблокировал урон: %d. Осталось хп: %d.\n", e.Name, blocked, e.Health)
}

func main() {
	myHero := Hero{Name: "GrigoDev", Health: 100, Attack: 20, Defense: 20}
	myEnemy := Enemy{Name: "Гусь", Health: 100, Attack: 50, Defense: 0}

	myEnemy.Attackayu()
	myHero.Defenseyu(myEnemy.Attack)

	fmt.Println("")

	myHero.Attackayu()
	myEnemy.Defenseyu(myHero.Attack)
}
