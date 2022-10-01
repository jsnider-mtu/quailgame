package items

type HeavyCrossbow struct {
}

func (h HeavyCrossbow) Slot() string {
    return "BothHands"
}

func (h HeavyCrossbow) Use() {
    // must be equipped to use
}

func (h HeavyCrossbow) Save() string {
    return "HeavyCrossbow"
}

func (h HeavyCrossbow) PrettyPrint() string {
    return "Heavy Crossbow"
}
