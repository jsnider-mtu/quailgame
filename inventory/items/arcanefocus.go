package items

type ArcaneFocus struct {
}

func (a ArcaneFocus) Slot() string {
    return "LeftHand"
}

func (a ArcaneFocus) Use() {
}

func (a ArcaneFocus) Save() string {
    return "ArcaneFocus"
}

func (a ArcaneFocus) PrettyPrint() string {
    return "Arcane Focus" 
}
