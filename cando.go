package cando

import (
	"errors"
)

const ErrorInvalidState = "error: fsm object is attempting to switch to an undefined state"
const ErrorStateSwitchInvalid = "error: fsm object is attempting to switch to a state that is not allowed to be switched to from the current one"

// State is a basic struct that allows you to define how a State functions.
type State struct {
	id       any
	OnEnter  func(prev, current *State, args ...any) // Enter is called when the state is entered from either no State, or a previous State.
	OnUpdate func(current *State, args ...any)       // Update is called whenever FSM.Update() is called and the state is active.
	OnExit   func(current, next *State, args ...any) // Exit is called when the state is exited as the FSM transitions from the current state to a new one.

	// Allowed is a function that determines if a State is allowed to be transitioned to from another
	// the FSM.Set() function.
	// If the function returns true, the state is switched; if not, the state remains the same as the original (sourceState).
	// If the function is not defined, then all state transitions are allowed.
	Allowed func(sourceState *State) bool

	// Allows you to define tags to identify and check for States.
	Tags []any
}

func (s *State) HasTag(tag any) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// FSM represents a Finite State Machine, which can have one State active at a time.
type FSM struct {
	currentState *State
	prevState    *State
	nextState    *State

	currentStateArgs []any
	nextStateArgs    []any

	States map[any]*State
}

// NewFSM creates a new FSM and returns it.
func NewFSM() *FSM {
	fsm := &FSM{}
	fsm.States = make(map[any]*State, 0)
	return fsm
}

// Update runs the Update() on the active State.
func (f *FSM) Update() {

	prevState := f.prevState
	nextState := f.nextState

	f.prevState = nil
	f.nextState = nil

	if prevState != nil {
		if prevState.OnExit != nil {
			prevState.OnExit(prevState, nextState, f.currentStateArgs...)
		}
	}

	if nextState != nil {
		if nextState.OnEnter != nil {
			nextState.OnEnter(prevState, nextState, f.nextStateArgs...)
		}
		f.currentState = nextState
	}

	f.currentStateArgs = f.nextStateArgs

	if f.currentState != nil && f.currentState.OnUpdate != nil {
		f.currentState.OnUpdate(f.currentState, f.currentStateArgs...)
	}

}

// Returns if the FSM is currently running a State with the given ID.
func (f *FSM) IsInStateWithID(stateID any) bool {
	return f.currentState != nil && f.currentState.id == stateID
}

// Returns if the FSM is currently running a State with all of the given tags.
func (f *FSM) IsInStateWithTags(tags ...any) bool {
	if f.currentState != nil {
		for _, tag := range tags {

			hasTag := false
			for _, t := range f.currentState.Tags {
				if t == tag {
					hasTag = true
					break
				}
			}
			if !hasTag {
				return false
			}

		}
		return true
	}
	return false
}

// Register registers a State with its id.
func (f *FSM) Register(id any, state *State) {
	state.id = id
	f.States[id] = state
}

// Unregister removes a State from the FSM using its id.
func (f *FSM) Unregister(id any) {
	delete(f.States, id)
}

// HasState returns if the FSM has a State associated with the id in its directory.
func (f *FSM) HasState(id any) bool {
	_, hasKey := f.States[id]
	return hasKey
}

// Allows you to set the current State assigned to the FSM.
// Any args passed will be accessible in the target State's OnEnter(), OnUpdate(), and OnExit() calls.
// After changing a State, it will call OnExit() on the previous State and
// OnEnter() on the next State on the next OnUpdate() call.
// If the state cannot be set, the function will return an error.
func (f *FSM) Set(toStateID any, args ...any) error {

	nextState, hasKey := f.States[toStateID]
	if !hasKey {
		return errors.New(ErrorInvalidState)
	}

	if nextState.Allowed != nil && !nextState.Allowed(f.currentState) {
		return errors.New(ErrorStateSwitchInvalid)
	}

	f.prevState = f.currentState
	f.nextState = nextState
	f.nextStateArgs = args

	return nil

}

// State returns the state machine's current State.
func (f *FSM) State() *State {
	return f.currentState
}
