package items

type LightCrossbow struct {
}

func (l *LightCrossbow) Slot() string {
    return "BothHands"
}

func (l *LightCrossbow) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (l *LightCrossbow) Save() string {
    return "LightCrossbow"
}

func (l *LightCrossbow) PrettyPrint() string {
    return "Light Crossbow"
}

func (l *LightCrossbow) Function() string {
    return "range"
}

func (l *LightCrossbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (l *LightCrossbow) Action() string {
    return ""
}

func (l *LightCrossbow) GetQuantity() int {
    return 1
}

func (l *LightCrossbow) GetRange() []float64 {
    return []float64{0, 0}
}
