package items

type HeavyCrossbow struct {
    Ammo int
}

func (h HeavyCrossbow) Slot() string {
    return "BothHands"
}

func (h HeavyCrossbow) Use() (string, []int) {
    return h.Action(), []int{}
}

func (h HeavyCrossbow) Save() string {
    return "HeavyCrossbow"
}

func (h HeavyCrossbow) PrettyPrint() string {
    return "Heavy Crossbow"
}

func (h HeavyCrossbow) Function() string {
    return "range-ammo-heavy-loading"
}

func (h HeavyCrossbow) Damage() (int, int, string) {
    return 1, 10, "piercing"
}

func (h HeavyCrossbow) Action() string {
    return ""
}

func (h HeavyCrossbow) GetQuantity() int {
    return h.Ammo
}

func (h HeavyCrossbow) GetRange() []float64 {
    return []float64{480.0, 1920.0}
}
