package items

type Shortbow struct {
}

func (s Shortbow) Slot() string {
    return "BothHands"
}

func (s Shortbow) Use() {
    // must be equipped to use
}

func (s Shortbow) Save() string {
    return "Shortbow"
}

func (s Shortbow) PrettyPrint() string {
    return "Shortbow"
}
