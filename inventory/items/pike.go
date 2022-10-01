package items

type Pike struct {
}

func (p Pike) Slot() string {
    return "BothHands"
}

func (p Pike) Use() {
    // must be equipped to use
}

func (p Pike) Save() string {
    return "Pike"
}

func (p Pike) PrettyPrint() string {
    return "Pike"
}
