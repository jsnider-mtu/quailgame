package items

type Spear struct {
}

func (s Spear) Slot() string {
    return "BothHands"
}

func (s Spear) Use() {
    // must be equipped to use
}

func (s Spear) Save() string {
    return "Spear"
}

func (s Spear) PrettyPrint() string {
    return "Spear"
}
