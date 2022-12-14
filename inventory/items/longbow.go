package items

type Longbow struct {
}

func (l *Longbow) Slot() string {
    return "BothHands"
}

func (l *Longbow) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (l *Longbow) Save() string {
    return "Longbow"
}

func (l *Longbow) PrettyPrint() string {
    return "Longbow"
}

func (l *Longbow) Function() string {
    return "range"
}

func (l *Longbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (l *Longbow) Action() string {
    return ""
}

func (l *Longbow) GetQuantity() int {
    return 1
}

func (l *Longbow) GetRange() []float64 {
    return []float64{0, 0}
}
