package items

type Handaxe struct {
}

func (h Handaxe) Slot() string {
    return "RightHand"
}

func (h Handaxe) Use() {
    // must be equipped to use
}

func (h Handaxe) Save() string {
    return "Handaxe"
}

func (h Handaxe) PrettyPrint() string {
    return "Handaxe"
}
