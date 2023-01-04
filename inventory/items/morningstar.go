package items

type Morningstar struct {
}

func (m *Morningstar) Slot() string {
    return "RightHand"
}

func (m *Morningstar) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (m *Morningstar) Save() string {
    return "Morningstar"
}

func (m *Morningstar) PrettyPrint() string {
    return "Morningstar"
}

func (m *Morningstar) Function() string {
    return "melee"
}

func (m *Morningstar) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (m *Morningstar) Action() string {
    return ""
}

func (m *Morningstar) GetQuantity() int {
    return 1
}

func (m *Morningstar) GetRange() []float64 {
    return []float64{0, 0}
}
