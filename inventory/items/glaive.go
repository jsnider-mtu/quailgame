package items

type Glaive struct {
}

func (g Glaive) Slot() string {
    return "RightHand"
}

func (g Glaive) Use() {
    // must be equipped to use
}

func (g Glaive) Save() string {
    return "Glaive"
}

func (g Glaive) PrettyPrint() string {
    return "Glaive"
}
