package items

type HandCrossbow struct {
}

func (h HandCrossbow) Slot() string {
    return "RightHand"
}

func (h HandCrossbow) Use() (string, []int) {
    return h.Action(), []int{}
}

func (h HandCrossbow) Save() string {
    return "HandCrossbow"
}

func (h HandCrossbow) PrettyPrint() string {
    return "Hand Crossbow"
}

func (h HandCrossbow) Function() string {
    return "range-ammo-light-loading"
}

func (h HandCrossbow) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (h HandCrossbow) Action() string {
    return ""
}

func (h HandCrossbow) GetQuantity() int {
    return 1
}

func (h HandCrossbow) GetRange() []float64 {
    return []float64{0, 0}
}
