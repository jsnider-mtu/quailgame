package items

type Pike struct {
}

func (p *Pike) Slot() string {
    return "BothHands"
}

func (p *Pike) Use() (string, []int) {
    return p.Action(), []int{}
}

func (p *Pike) Save() string {
    return "Pike"
}

func (p *Pike) PrettyPrint() string {
    return "Pike"
}

func (p *Pike) Function() string {
    return "melee-heavy-reach"
}

func (p *Pike) Damage() (int, int, string) {
    return 1, 10, "piercing"
}

func (p *Pike) Action() string {
    return ""
}

func (p *Pike) GetQuantity() int {
    return 1
}

func (p *Pike) GetRange() []float64 {
    return []float64{24.0, 24.0}
}
