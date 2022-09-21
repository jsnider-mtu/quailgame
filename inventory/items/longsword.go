package items

type Longsword struct {
}

func (l Longsword) Slot() string {
    return "RightHand"
}

func (l Longsword) Use() {
    // must be equipped to use
}

func (l Longsword) Save() string {
    return "Longsword"
}
