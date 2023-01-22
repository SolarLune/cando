package cando

import "fmt"

// IState is a basic interface that works like you would expect.
type IState interface {
	Enter()
	Update()
	Exit()
}

// TODO: Add something to define and return if states can enter into other ones.

// FSM represents a Finite State Machine, which can have one State active at a time.
type FSM struct {
	CurrentState string
	States       map[string]IState
}

// NewFSM creates a new FSM and returns it.
func NewFSM() *FSM {
	fsm := &FSM{}
	fsm.States = make(map[string]IState, 0)
	return fsm
}

// Update runs the Update() on the active State.
func (f *FSM) Update() {
	if f.CurrentState != "" {
		f.States[f.CurrentState].Update()
	} else {
		fmt.Println("Update() called on FSM without active state.")
	}
}

// Register registers a State with its name.
func (f *FSM) Register(name string, state IState) {
	f.States[name] = state
}

// Unregister removes a State from the FSM using its name.
func (f *FSM) Unregister(name string) {
	delete(f.States, name)
}

// HasState returns if the FSM has a State associated with the name in its directory.
func (f *FSM) HasState(name string) bool {
	_, hasKey := f.States[name]
	return hasKey
}

// Change allows you to change the current, "main" State assigned to the FSM. If you run Change(), it will call
// Exit() on the previous State and Enter() on the next State.
func (f *FSM) Change(stateName string) {

	if f.CurrentState != "" {
		f.States[f.CurrentState].Exit()
	}

	_, hasKey := f.States[stateName]
	if !hasKey {
		panic("Error: FSM object is attempting to switch to an invalid / undefined state: " + stateName)
	}

	f.CurrentState = stateName

	if f.CurrentState != "" && f.States[f.CurrentState].Enter != nil {
		f.States[f.CurrentState].Enter()
	}

}
