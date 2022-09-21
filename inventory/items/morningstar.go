package items

type Morningstar struct {
}

func (m Morningstar) Slot() string {
    return "RightHand"
}

func (m Morningstar) Use() {
    // must be equipped to use
}

func (m Morningstar) Save() string {
    return "Morningstar"
}
