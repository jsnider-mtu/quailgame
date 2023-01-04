package items

type Maul struct {
}

func (m *Maul) Slot() string {
    return "BothHands"
}

func (m *Maul) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (m *Maul) Save() string {
    return "Maul"
}

func (m *Maul) PrettyPrint() string {
    return "Maul"
}

func (m *Maul) Function() string {
    return "melee"
}

func (m *Maul) Damage() (int, int, string) {
    return 2, 6, "bludgeoning"
}

func (m *Maul) Action() string {
    return ""
}

func (m *Maul) GetQuantity() int {
    return 1
}

func (m *Maul) GetRange() []float64 {
    return []float64{0, 0}
}
