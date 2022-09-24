package items

type Lance struct {
}

func (l Lance) Slot() string {
    return "RightHand"
}

func (l Lance) Use() {
    // must be equipped to use
}

func (l Lance) Save() string {
    return "Lance"
}

func (l Lance) PrettyPrint() string {
    return "Lance"
}
