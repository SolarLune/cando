package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/solarlune/cando"
)

type Human struct {
	FSM    *cando.FSM
	Hunger int
	Thirst int
}

func main() {

	// Make a new Human object
	human := &Human{
		Hunger: 100,
		FSM:    cando.NewFSM(),
	}

	turnsInState := 0

	human.FSM.Register("idle",

		&cando.State{

			OnEnter: func(prev *cando.State, current *cando.State, args ...any) { turnsInState = 0 },

			OnUpdate: func(current *cando.State, args ...any) {

				human.Hunger -= 10

				turnsInState++

				fmt.Println("IdleState OnUpdate: Idle Turn", turnsInState, ": My hunger level is at: ", human.Hunger)

				if human.Hunger <= 20 {
					fmt.Println("IdleState OnUpdate: I've reached my limit; I'm getting something to eat.")
					human.FSM.Set("search")
				}

			},
		},
	)

	human.FSM.Register("search",

		&cando.State{

			OnEnter: func(prev *cando.State, current *cando.State, args ...any) { turnsInState = 0 },

			OnUpdate: func(current *cando.State, args ...any) {

				turnsInState++

				if rand.Float32() < 0.2 {
					fmt.Println("SearchState OnUpdate: Search Turn", turnsInState, ": Ah, found something to eat!")
					human.FSM.Set("eating")
				} else {
					fmt.Println("SearchState OnUpdate: Search Turn", turnsInState, ": Hmm... I searched, but there wasn't anything to eat.")
				}

			},
		},
	)

	human.FSM.Register("eating",

		&cando.State{

			OnEnter: func(prev *cando.State, current *cando.State, args ...any) {
				turnsInState = 0
				fmt.Println("EatingState OnEnter: Finally, some good grub!")
			},

			OnUpdate: func(current *cando.State, args ...any) {
				turnsInState++
				human.Hunger += 10
				fmt.Println("EatingState OnUpdate: Eating Turn", turnsInState, ": *Chomp* *Smack* : ", human.Hunger)
				if human.Hunger >= 100 {
					human.FSM.Set("idle")
				}
			},

			OnExit: func(current *cando.State, next *cando.State, args ...any) {
				fmt.Println("EatingState OnExit: Phew, that was good. Back to doing nothing.")
			},
		})

	human.FSM.Set("idle")

	for {
		human.FSM.Update()
		time.Sleep(time.Millisecond * 500)
	}

}
