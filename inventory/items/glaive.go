package items

type Glaive struct {
}

func (g *Glaive) Slot() string {
    return "BothHands"
}

func (g *Glaive) Use() (string, []int) {
    return g.Action(), []int{}
}

func (g *Glaive) Save() string {
    return "Glaive"
}

func (g *Glaive) PrettyPrint() string {
    return "Glaive"
}

func (g *Glaive) Function() string {
    return "melee-heavy-reach"
}

func (g *Glaive) Damage() (int, int, string) {
    return 1, 10, "slashing"
}

func (g *Glaive) Action() string {
    return ""
}

func (g *Glaive) GetQuantity() int {
    return 1
}

func (g *Glaive) GetRange() []float64 {
    return []float64{0, 0}
}
