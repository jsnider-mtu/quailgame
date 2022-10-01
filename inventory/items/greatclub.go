package items

type Greatclub struct {
}

func (g Greatclub) Slot() string {
    return "RightHand"
}

func (g Greatclub) Use() {
    // must be equipped to use
}

func (g Greatclub) Save() string {
    return "Greatclub"
}

func (g Greatclub) PrettyPrint() string {
    return "Greatclub"
}
