package items

type Rapier struct {
}

func (r Rapier) Slot() string {
    return "RightHand"
}

func (r Rapier) Use() {
    // must be equipped to use
}

func (r Rapier) Save() string {
    return "Rapier"
}
