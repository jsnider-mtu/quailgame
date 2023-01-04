package items

type WarPick struct {
}

func (w WarPick) Slot() string {
    return "RightHand"
}

func (w WarPick) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (w WarPick) Save() string {
    return "WarPick"
}

func (w WarPick) PrettyPrint() string {
    return "War Pick"
}

func (w WarPick) Function() string {
    return "melee"
}

func (w WarPick) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (w WarPick) Action() string {
    return ""
}

func (w WarPick) GetQuantity() int {
    return 1
}

func (w WarPick) GetRange() []float64 {
    return []float64{0, 0}
}
