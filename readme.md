
# gofsm

Gofsm is a simple finite state machine for Golang. 

# Why

I created it because I wanted something super-duper simple for game development, and thought the existing FSMs for Go were a bit over-engineered for my needs.

# Installation

`go get github.com/SolarLune/gofsm` 

# Usage

```go

type Human struct {
    FSM *gofsm.FSM
}

func StateAwake() {

    fmt.Println("I'm awake! What a nice day.")

}

func StateSleep() {

    fmt.Println("ZZZ")

}

func main() {

    human := Human{}
    
    human.FSM = gofsm.NewFSM()

    // Here, we register a new State for the FSM, which has a pointer to the 
    // StateAwake function as the State's Update function.
    // This means that whenever we call Update() on the FSM and the FSM's 
    // current State is the "awake" State, it'll run StateAwake().
    
    human.FSM.Register("awake", gofsm.State{ Update: StateAwake } )

    human.FSM.Register("sleeping", gofsm.State{ Update: StateSleep } )
    
    human.FSM.Change("sleeping")

    human.FSM.Update() // <-- Prints "ZZZ"

}

```

You can also create a data structure to contain data that is linked to a State, and even pass that struct's methods to the State.

```go

type Eating struct {

    Fullness int32
    State gofsm.State

}

func (a *Eating) Update() {

    if a.Fullness < 100 {
        a.Fullness++
    }

}

func NewEating() Eating {

    a := Eating{}
    a.State = gofsm.State{ Update: a.Update }
    return a

}

```

Then you could pass a newly created Eating struct's `State` field into an FSM to update, and it would increment the Eating struct's Fullness value.
