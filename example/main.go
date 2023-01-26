package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/solarlune/cando"
)

// -------------------------------

func IdleState(human *Human) cando.State {

	turnsInState := 0

	return cando.State{

		Enter: func() { turnsInState = 0 },

		Update: func() {

			human.Hunger -= 10

			turnsInState++

			fmt.Println("Idle Turn", turnsInState, ": My hunger level is at: ", human.Hunger)

			if human.Hunger <= 20 {
				fmt.Println("I've reached my limit; I'm getting something to eat.")
				human.FSM.Change("search")
			}

		},
	}

}

// -------------------------------

func SearchState(human *Human) cando.State {

	turnsInState := 0

	return cando.State{

		Enter: func() { turnsInState = 0 },

		Update: func() {

			turnsInState++

			if rand.Float32() < 0.2 {
				fmt.Println("Search Turn", turnsInState, ": Ah, found something to eat!")
				human.FSM.Change("eating")
			} else {
				fmt.Println("Search Turn", turnsInState, ": Hmm... I searched, but there wasn't anything to eat.")
			}

		},
	}

}

// -------------------------------

func EatingState(human *Human) cando.State {

	turnsInState := 0

	return cando.State{

		Enter: func() {
			turnsInState = 0
			fmt.Println("Finally, some good grub!")
		},

		Update: func() {
			turnsInState++
			human.Hunger += 10
			fmt.Println("Eating Turn", turnsInState, ": *Chomp* *Smack* : ", human.Hunger)
			if human.Hunger >= 100 {
				human.FSM.Change("idle")
			}
		},

		Exit: func() { fmt.Println("Phew, that was good. Back to doing nothing.") },
	}

}

// -------------------------------

type Human struct {
	FSM    *cando.FSM
	Hunger int
	Thirst int
}

func NewHuman() *Human {

	human := &Human{
		Hunger: 100,
		FSM:    cando.NewFSM(),
	}

	human.FSM.Register("idle", IdleState(human))
	human.FSM.Register("search", SearchState(human))
	human.FSM.Register("eating", EatingState(human))
	human.FSM.Change("idle")

	return human
}

func main() {

	human := NewHuman()
	for {
		human.FSM.Update()
		time.Sleep(time.Millisecond * 500)
	}

}
