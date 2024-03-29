package items

type Mace struct {
}

func (m *Mace) Slot() string {
    return "RightHand"
}

func (m *Mace) Use() (string, []int) {
    return m.Action(), []int{}
}

func (m *Mace) Save() string {
    return "Mace"
}

func (m *Mace) PrettyPrint() string {
    return "Mace"
}

func (m *Mace) Function() string {
    return "melee"
}

func (m *Mace) Damage() (int, int, string) {
    return 1, 6, "bludgeoning"
}

func (m *Mace) Action() string {
    return ""
}

func (m *Mace) GetQuantity() int {
    return 1
}

func (m *Mace) GetRange() []float64 {
    return []float64{0, 0}
}
