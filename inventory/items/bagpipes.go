package items

type Bagpipes struct {
}

func (b *Bagpipes) Slot() string {
    return "BothHands"
}

func (b *Bagpipes) Use() (string, []int) {
    return b.Action(), []int{}
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
    return "playmusic"
}

func (b *Bagpipes) GetQuantity() int {
    return 1
}

func (b *Bagpipes) GetRange() []float64 {
    return []float64{0, 0}
}
