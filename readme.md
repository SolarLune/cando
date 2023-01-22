
# CanDo 💪

CanDo is a simple finite state machine for Golang. 

# Why

I created it because I wanted something super-duper simple for game development, and thought the existing FSMs for Go were a bit over-engineered for my needs (where I know which states I'm switching into for an object and already know that the states are valid, for example).

# Installation

`go get github.com/solarlune/cando` 

# Usage

```go

type Eating struct {
    Hunger int32
}

func (e *Eating) Enter() {
    e.Hunger = 100
}

func (e *Eating) Update() {
    if e.Hunger > 0 {
        e.Hunger--
    }
}

func (e *Eating) Exit() {
    fmt.Println("*Burp*! All done!")
}

func main() {
    
    fsm := cando.NewFSM()
    fsm.Register("eating", Eating{})
    fsm.Change("eating") // Eating.Hunger == 100 now, and will go down by 1 each time we call fsm.Update().

    fsm.Update() // Eating.Hunger == 99 now.

    // If we switch from the "eating" state to another one, then it will call Finish() on the Eating struct.

}

// That's also it!

```

## Didn't this used to be different?

Yeah, I updated it; now it's more conventional and generally easier to deal with.

## To-do

- [ ] Add some method of indicating which states may be passed into from other ones
- [ ] Add decision trees
