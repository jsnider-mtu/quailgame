package items

type Trident struct {
}

func (t Trident) Slot() string {
    return "RightHand"
}

func (t Trident) Use() {
    // must be equipped to use
}

func (t Trident) Save() string {
    return "Trident"
}

func (t Trident) PrettyPrint() string {
    return "Trident"
}
