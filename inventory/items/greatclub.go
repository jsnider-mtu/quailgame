package items

type Greatclub struct {
}

func (g *Greatclub) Slot() string {
    return "BothHands"
}

func (g *Greatclub) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
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
