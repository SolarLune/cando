
# CanDo 💪

CanDo is a simple finite state machine for Golang. 

# Why

I created it because I wanted something super-duper simple for game development, and thought the existing FSMs for Go were a bit over-engineered for my needs (primarily my circumstances where I know which states I'm switching into for an object and already know that the states are valid, for example). 

I also choose eschew the traditional route of creating an FSM using a state interface that you implement for each state. 

While this works fine, it necessitates boilerplate code, particularly if your states have no Enter or Exit functions, for example. It _also_ forces you to define states on the package-level as interface-implementing structs, naturally.

Instead, I opt to take the following approach: a State is a struct that contains pointers to Enter(), Exit(), and Update() functions. A *cando.FSM (Finite State Machine) facilitates switching from, to, and updating the active State. If a State doesn't define an `Enter()`, `Exit()`, or `Update()` function, then the function isn't called - simple as that.

This allows you to more flexibly and simply define states.

# Installation

`go get github.com/solarlune/cando` 

# Usage

```go

// A cando.State is a struct that contains optional Enter, Update, and Exit function pointers.
// This being the case, the State can be constructed as you wish - from a standard
// constructor function, as function literals, as pointers to existing functions on a 
// struct, etc.

// cando.FSM is the Finite State Machine that handles switching to and from, and
// updating the currently active State.

// State.Enter() is called when the State is switched to.
// State.Update() is called however often FSM.Update() is called.
// State.Exit() is called when the State is switched from.

// In this example, we create a constructor function to return an "Eating State".
func EatingState() cando.State {

    // Locally-scoped variables can easily serve as state fields.
    hunger := 0

    return cando.State{
        
        Enter: func() { hunger = 100 },
        Update: func() {
            if hunger > 0 {
                hunger--
            }
        },
        Exit: func() { fmt.Println("*Burp*! All done!") },

    }

}

func main() {
    
    fsm := cando.NewFSM() // Create an FSM.
    fsm.Register("eating", EatingState()) // Register the eating state.
    fsm.Change("eating") // Switch to the eating state; calls State.Enter(), if defined

    fsm.Update() // EatingState.Hunger == 99 now.

    // If we switch from the "eating" state to another one, then it will
    // call Exit() on the Eating struct.

}

// That's also it!

```

## To-do

- [ ] Add some method of indicating which states may be passed into from other ones
- [ ] Add decision trees?
