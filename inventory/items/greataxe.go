package items

type Greataxe struct {
}

func (g *Greataxe) Slot() string {
    return "BothHands"
}

func (g *Greataxe) Use() (string, []int) {
    return g.Action(), []int{}
}

func (g *Greataxe) Save() string {
    return "Greataxe"
}

func (g *Greataxe) PrettyPrint() string {
    return "Greataxe"
}

func (g *Greataxe) Function() string {
    return "melee-heavy"
}

func (g *Greataxe) Damage() (int, int, string) {
    return 1, 12, "slashing"
}

func (g *Greataxe) Action() string {
    return ""
}

func (g *Greataxe) GetQuantity() int {
    return 1
}

func (g *Greataxe) GetRange() []float64 {
    return []float64{0, 0}
}
