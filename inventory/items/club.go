package items

type Club struct {
}

func (c Club) Slot() string {
    return "RightHand"
}

func (c Club) Use() {
    // must be equipped to use
}

func (c Club) Save() string {
    return "Club"
}
