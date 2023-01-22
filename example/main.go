package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/solarlune/cando"
)

// -------------------------------

type IdleState struct {
	Human *Human
}

func (state IdleState) Enter() {}
func (state IdleState) Update() {

	state.Human.Hunger -= 10

	hunger := state.Human.Hunger

	fmt.Println("My hunger level is at: ", state.Human.Hunger)

	if hunger <= 20 {
		fmt.Println("I've reached my limit; I'm getting something to eat.")
		state.Human.FSM.Change("search")
	}

}
func (state IdleState) Exit() {}

// -------------------------------

type SearchState struct {
	Human *Human
}

func (state SearchState) Enter() {}
func (state SearchState) Update() {
	if rand.Float32() < 0.2 {
		fmt.Println("Ah, found something to eat!")
		state.Human.FSM.Change("eating")
	} else {
		fmt.Println("Hmm... I searched, but there wasn't anything to eat.")
	}
}
func (state SearchState) Exit() {}

// -------------------------------

type EatingState struct {
	Human *Human
}

func (state EatingState) Enter() { fmt.Println("Finally, some good grub!") }
func (state EatingState) Update() {
	state.Human.Hunger += 10
	fmt.Println("*Chomp* *Smack* : ", state.Human.Hunger)
	if state.Human.Hunger >= 100 {
		state.Human.FSM.Change("idle")
	}
}
func (state EatingState) Exit() { fmt.Println("Phew, that was good. Back to doing nothing.") }

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

	human.FSM.Register("idle", IdleState{Human: human})
	human.FSM.Register("search", SearchState{Human: human})
	human.FSM.Register("eating", EatingState{Human: human})
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
