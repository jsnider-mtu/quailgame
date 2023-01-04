package items

type Warhammer struct {
}

func (w *Warhammer) Slot() string {
    return "BothHands"
}

func (w *Warhammer) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (w *Warhammer) Save() string {
    return "Warhammer"
}

func (w *Warhammer) PrettyPrint() string {
    return "Warhammer"
}

func (w *Warhammer) Function() string {
    return "melee"
}

func (w *Warhammer) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}

func (w *Warhammer) Action() string {
    return ""
}

func (w *Warhammer) GetQuantity() int {
    return 1
}

func (w *Warhammer) GetRange() []float64 {
    return []float64{0, 0}
}
