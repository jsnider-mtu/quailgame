package items

type Greatclub struct {
}

func (g *Greatclub) Slot() string {
    return "BothHands"
}

func (g *Greatclub) Use() (string, []int) {
    return g.Action(), []int{}
}

func (g *Greatclub) Save() string {
    return "Greatclub"
}

func (g *Greatclub) PrettyPrint() string {
    return "Greatclub"
}

func (g *Greatclub) Function() string {
    return "melee"
}

func (g *Greatclub) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}

func (g *Greatclub) Action() string {
    return ""
}

func (g *Greatclub) GetQuantity() int {
    return 1
}

func (g *Greatclub) GetRange() []float64 {
    return []float64{0, 0}
}
