package gofsm

import "fmt"

// State is a basic struct that implements the State interface.
type State struct {
	Enter  func()
	Update func()
	Exit   func()
}

// FSM represents a Finite State Machine, which can have one State active at a time.
type FSM struct {
	State          string
	StateDirectory map[string]State
}

// NewFSM creates a new FSM and returns it.
func NewFSM() *FSM {
	fsm := &FSM{}
	fsm.StateDirectory = make(map[string]State, 0)
	return fsm
}

// Update runs the Update() on the active State.
func (f *FSM) Update() {
	if f.State != "" && f.StateDirectory[f.State].Update != nil {
		f.StateDirectory[f.State].Update()
	} else {
		fmt.Println("Update() called on FSM without active state.")
	}
}

// Register registers a State with its name.
func (f *FSM) Register(name string, state State) {
	f.StateDirectory[name] = state
}

// Unregister removes a State from the FSM using its name.
func (f *FSM) Unregister(name string) {
	delete(f.StateDirectory, name)
}

// HasState returns if the FSM has a State associated with the name in its directory.
func (f *FSM) HasState(name string) bool {
	_, hasKey := f.StateDirectory[name]
	return hasKey
}

// Change allows you to change the current, "main" State assigned to the FSM. If you run Change(), it will call
// Exit() on the previous State and Enter() on the next State.
func (f *FSM) Change(stateName string) {

	if f.State != "" && f.StateDirectory[f.State].Exit != nil {
		f.StateDirectory[f.State].Exit()
	}

	_, hasKey := f.StateDirectory[stateName]
	if !hasKey {
		fmt.Println("FSM object", f, "has no state", stateName)
		panic("Error!")
	}

	f.State = stateName

	if f.State != "" && f.StateDirectory[f.State].Enter != nil {
		f.StateDirectory[f.State].Enter()
	}

}
