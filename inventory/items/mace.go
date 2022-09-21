package items

type Mace struct {
}

func (m Mace) Slot() string {
    return "RightHand"
}

func (m Mace) Use() {
    // must be equipped to use
}

func (m Mace) Save() string {
    return "Mace"
}
