package items

type LeatherArmor struct {
}

func (l *LeatherArmor) Slot() string {
    return "Armor"
}

func (l *LeatherArmor) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *LeatherArmor) Save() string {
    return "LeatherArmor"
}

func (l *LeatherArmor) PrettyPrint() string {
    return "Leather Armor"
}

func (l *LeatherArmor) Function() string {
    return "armor"
}

func (l *LeatherArmor) Damage() (int, int, string) {
    return 0, 0, ""
}

func (l *LeatherArmor) Action() string {
    return ""
}

func (l *LeatherArmor) GetQuantity() int {
    return 1
}

func (l *LeatherArmor) GetRange() []float64 {
    return []float64{0, 0}
}
