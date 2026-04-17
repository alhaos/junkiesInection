package main

import (
	"fmt"

	"github.com/alhaos/junkiesInection/internal/model"
)

const count = 1_000_000

func main() {

	var sum int

	for _ = range count {
		t := model.NewDefaultTrial()
		i := 0
		for ; !t.IsAllInfected(); i++ {
			t.ShootUpJunkies()
		}
		sum += i
	}

	fmt.Printf("avg days: %d", sum/count)
}
