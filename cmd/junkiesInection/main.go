package main

import (
	"fmt"

	"github.com/alhaos/junkiesInection/internal/model"
)

func main() {

	for n := range 100 {
		t := model.NewDefaultTrial()
		i := 0
		for ; !t.IsAllInfected(); i++ {
			t.ShootUpJunkies()
		}
		fmt.Printf("case: [%d], days:[%d]\r\n", n, i)
	}
}
