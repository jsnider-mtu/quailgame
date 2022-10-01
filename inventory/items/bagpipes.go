package items

type Bagpipes struct {
}

func (b Bagpipes) Slot() string {
    return "BothHands"
}

func (b Bagpipes) Use() {
}

func (b Bagpipes) Save() string {
    return "Bagpipes"
}

func (b Bagpipes) PrettyPrint() string {
    return "Bagpipes" 
}
