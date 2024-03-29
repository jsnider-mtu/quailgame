package items

type Lyre struct {
}

func (l *Lyre) Slot() string {
    return "BothHands"
}

func (l *Lyre) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *Lyre) Save() string {
    return "Lyre"
}

func (l *Lyre) PrettyPrint() string {
    return "Lyre"
}

func (l *Lyre) Function() string {
    return "instrument"
}

func (l *Lyre) Damage() (int, int, string) {
    return 0, 0, ""
}

func (l *Lyre) Action() string {
    return "playmusic"
}

func (l *Lyre) GetQuantity() int {
    return 1
}

func (l *Lyre) GetRange() []float64 {
    return []float64{0, 0}
}
