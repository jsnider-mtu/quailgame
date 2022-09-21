package items

type Greataxe struct {
}

func (g Greataxe) Slot() string {
    return "RightHand"
}

func (g Greataxe) Use() {
    // must be equipped to use
}

func (g Greataxe) Save() string {
    return "Greataxe"
}
