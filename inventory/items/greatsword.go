package items

type Greatsword struct {
}

func (g *Greatsword) Slot() string {
    return "BothHands"
}

func (g *Greatsword) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (g *Greatsword) Save() string {
    return "Greatsword"
}

func (g *Greatsword) PrettyPrint() string {
    return "Greatsword"
}

func (g *Greatsword) Function() string {
    return "melee"
}

func (g *Greatsword) Damage() (int, int, string) {
    return 2, 6, "slashing"
}

func (g *Greatsword) Action() string {
    return ""
}

func (g *Greatsword) GetQuantity() int {
    return 1
}

func (g *Greatsword) GetRange() []float64 {
    return []float64{0, 0}
}
