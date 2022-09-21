package items

type Javelin struct {
}

func (j Javelin) Slot() string {
    return "RightHand"
}

func (j Javelin) Use() {
    // must be equipped to use
}

func (j Javelin) Save() string {
    return "Javelin"
}
