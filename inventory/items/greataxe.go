package items

type Greataxe struct {
}

func (g Greataxe) Slot() string {
    return "BothHands"
}

func (g Greataxe) Use() {
    // must be equipped to use
}

func (g Greataxe) Save() string {
    return "Greataxe"
}

func (g Greataxe) PrettyPrint() string {
    return "Greataxe"
}
