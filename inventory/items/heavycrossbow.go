package items

type HeavyCrossbow struct {
}

func (h HeavyCrossbow) Slot() string {
    return "BothHands"
}

func (h HeavyCrossbow) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (h HeavyCrossbow) Save() string {
    return "HeavyCrossbow"
}

func (h HeavyCrossbow) PrettyPrint() string {
    return "Heavy Crossbow"
}

func (h HeavyCrossbow) Function() string {
    return "range"
}

func (h HeavyCrossbow) Damage() (int, int, string) {
    return 1, 10, "piercing"
}

func (h HeavyCrossbow) Action() string {
    return ""
}

func (h HeavyCrossbow) GetQuantity() int {
    return 1
}
