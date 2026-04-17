package model

import (
	"fmt"
	"math/rand"
	"slices"
)

type Trial struct {
	Junkies []Junkie
}

func NewDefaultTrial() *Trial {

	t := Trial{}

	t.Junkies = append(t.Junkies, Junkie{Infected: true, ID: 0})

	for i := range 9 {
		t.Junkies = append(t.Junkies, Junkie{Infected: false, ID: i + 1})
	}

	return &t
}

func (t *Trial) ShootUpJunkies() {

	rand.Shuffle(len(t.Junkies), func(i, j int) {
		t.Junkies[i], t.Junkies[j] = t.Junkies[j], t.Junkies[i]
	})

	isNeedleInfected := false

	for i, junkie := range t.Junkies {

		if junkie.Infected {
			isNeedleInfected = true
		}

		if isNeedleInfected {
			t.Junkies[i].Infected = t.Junkies[i].Infected || (rand.Intn(100) < 2)
		}
	}
}

func (t *Trial) PrintState() {

	for _, j := range t.Junkies {
		fmt.Printf("%d: %s\r\n", j.ID, func(infected bool) string {
			if infected {
				return "infected"
			}
			return "clear"
		}(j.Infected))
	}
}

func (t *Trial) SortByID() {
	slices.SortFunc(t.Junkies, func(a, b Junkie) int {
		return a.ID - b.ID
	})
}

func (t *Trial) IsAllInfected() bool {

	var result bool

	result = true

	for _, j := range t.Junkies {
		result = result && j.Infected
	}

	return result
}
