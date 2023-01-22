
# CanDo 💪

CanDo is a simple finite state machine for Golang. 

# Why

I created it because I wanted something super-duper simple for game development, and thought the existing FSMs for Go were a bit over-engineered for my needs (where I know which states I'm switching into for an object and already know that the states are valid, for example).

# Installation

`go get github.com/SolarLune/gofsm` 

# Usage

```go

func StateAwake() {
    fmt.Println("I'm awake! What a nice day.")
}

func StateSleep() {
    fmt.Println("ZZZ")
}

func main() {

    human := gofsm.NewFSM()
    
    // Here, we register a new State for the FSM, called "awake". It has a 
    // pointer to the StateAwake function as the State's Update function call.

    human.Register("awake", gofsm.State{ Update: StateAwake } )

    // This means that whenever we call Update() on the FSM and the FSM's 
    // current State is the "awake" State, it'll run StateAwake().

    human.Register("sleeping", gofsm.State{ Update: StateSleep } )
    
    // Here, we change the active state to "sleeping".
    human.Change("sleeping")

    human.Update() // <-- Prints "ZZZ"

}

// That's it!

```

You can also create a data structure to contain data that is linked to a State, and even pass that struct's methods to the State (which doesn't need to exist outside of the FSM, since it's essentially just a boilerplate for a collection of function pointers).

```go

type Eating struct {
    Hunger int32
}

func (e *Eating) Begin() {
    e.Hunger = 100
}

func (e *Eating) Eat() {
    if e.Hunger > 0 {
        e.Hunger--
    }
}

func (e *Eating) Finish() {
    fmt.Println("*Burp*! All done!")
}

func main() {

    es := Eating{}

    fsm := gofsm.NewFSM()
    fsm.Register("eating", gofsm.State { Enter: es.Begin, Update: es.Eat, Exit: es.Finish })
    fsm.Change("eating") // eating.Hunger == 100 now.

    // If we switch from the "eating" state to another one, then it will call Finish() on the eating struct.

}

// That's also it!

```

## Why didn't you go (haha) the "correct" route and have the State be an interface that could be implemented by any fulfilling struct?

For simplicity, States are hard-coded and you simply override the functions. Not all functions need to be implemented, as by default, States don't call Enter, Update, or Exit functions unless they're defined. This cuts down on boilerplate code considerably.

## To-do

- [ ] Add some method of indicating which states may be passed into from other ones
- [ ] Add decision trees
