package items

type Glaive struct {
}

func (g *Glaive) Slot() string {
    return "BothHands"
}

func (g *Glaive) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (g *Glaive) Save() string {
    return "Glaive"
}

func (g *Glaive) PrettyPrint() string {
    return "Glaive"
}

func (g *Glaive) Function() string {
    return "melee"
}

func (g *Glaive) Damage() (int, int, string) {
    return 1, 10, "slashing"
}

func (g *Glaive) Action() string {
    return ""
}
