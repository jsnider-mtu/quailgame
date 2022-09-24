package items

type Blowgun struct {
}

func (b Blowgun) Slot() string {
    return "RightHand"
}

func (b Blowgun) Use() {
    // must be equipped to use
}

func (b Blowgun) Save() string {
    return "Blowgun"
}

func (b Blowgun) PrettyPrint() string {
    return "Blowgun"
}
