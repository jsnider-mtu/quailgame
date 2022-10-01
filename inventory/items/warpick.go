package items

type Warpick struct {
}

func (w Warpick) Slot() string {
    return "RightHand"
}

func (w Warpick) Use() {
    // must be equipped to use
}

func (w Warpick) Save() string {
    return "Warpick"
}

func (w Warpick) PrettyPrint() string {
    return "Warpick"
}
