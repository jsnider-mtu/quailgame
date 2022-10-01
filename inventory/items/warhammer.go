package items

type Warhammer struct {
}

func (w Warhammer) Slot() string {
    return "BothHands"
}

func (w Warhammer) Use() {
    // must be equipped to use
}

func (w Warhammer) Save() string {
    return "Warhammer"
}

func (w Warhammer) PrettyPrint() string {
    return "Warhammer"
}
