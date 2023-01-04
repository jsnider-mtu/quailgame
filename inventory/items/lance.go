package items

type Lance struct {
}

func (l *Lance) Slot() string {
    return "RightHand"
}

func (l *Lance) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *Lance) Save() string {
    return "Lance"
}

func (l *Lance) PrettyPrint() string {
    return "Lance"
}

func (l *Lance) Function() string {
    return "melee-reach-special"
}

func (l *Lance) Damage() (int, int, string) {
    return 1, 12, "piercing"
}

func (l *Lance) Action() string {
    return ""
}

func (l *Lance) GetQuantity() int {
    return 1
}

func (l *Lance) GetRange() []float64 {
    return []float64{0, 0}
}
