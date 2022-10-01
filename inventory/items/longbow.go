package items

type Longbow struct {
}

func (l Longbow) Slot() string {
    return "RightHand"
}

func (l Longbow) Use() {
    // must be equipped to use
}

func (l Longbow) Save() string {
    return "Longbow"
}

func (l Longbow) PrettyPrint() string {
    return "Longbow"
}
