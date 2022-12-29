package items

type Bagpipes struct {
}

func (b *Bagpipes) Slot() string {
    return "BothHands"
}

func (b *Bagpipes) Use() (string, []int) {
    return "", []int{}
}

func (b *Bagpipes) Save() string {
    return "Bagpipes"
}

func (b *Bagpipes) PrettyPrint() string {
    return "Bagpipes" 
}

func (b *Bagpipes) Function() string {
    return "instrument"
}

func (b *Bagpipes) Damage() (int, int, string) {
    return 0, 0, ""
}

func (b *Bagpipes) Action() string {
    return ""
}

func (b *Bagpipes) GetQuantity() int {
    return 1
}
