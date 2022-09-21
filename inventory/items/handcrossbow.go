package items

type Handcrossbow struct {
}

func (h Handcrossbow) Slot() string {
    return "RightHand"
}

func (h Handcrossbow) Use() {
    // must be equipped to use
}

func (h Handcrossbow) Save() string {
    return "Handcrossbow"
}
