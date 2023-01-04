package items

type Longsword struct {
}

func (l *Longsword) Slot() string {
    return "RightHand"
}

func (l *Longsword) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (l *Longsword) Save() string {
    return "Longsword"
}

func (l *Longsword) PrettyPrint() string {
    return "Longsword"
}

func (l *Longsword) Function() string {
    return "melee"
}

func (l *Longsword) Damage() (int, int, string) {
    return 1, 8, "slashing"
}

func (l *Longsword) Action() string {
    return ""
}

func (l *Longsword) GetQuantity() int {
    return 1
}

func (l *Longsword) GetRange() []float64 {
    return []float64{0, 0}
}
