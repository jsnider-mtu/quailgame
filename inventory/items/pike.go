package items

type Pike struct {
}

func (p Pike) Slot() string {
    return "RightHand"
}

func (p Pike) Use() {
    // must be equipped to use
}

func (p Pike) Save() string {
    return "Pike"
}
