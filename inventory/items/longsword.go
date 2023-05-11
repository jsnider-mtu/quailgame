package items

type Longsword struct {
    Currslot string
}

func (l *Longsword) GetCurrSlot() string {
    return l.Currslot
}

func (l *Longsword) SwitchSlots() {
    if l.Currslot == "BothHands" {
        l.Currslot = "RightHand"
    } else if l.Currslot == "RightHand" {
        l.Currslot = "BothHands"
    }
    return
}

func (l *Longsword) Slot() string {
    return "Versatile"
}

func (l *Longsword) Use() (string, []int) {
    return l.Action(), []int{1, 10}
}

func (l *Longsword) Save() string {
    return "Longsword"
}

func (l *Longsword) PrettyPrint() string {
    return "Longsword"
}

func (l *Longsword) Function() string {
    return "melee"
}

func (l *Longsword) Damage() (int, int, string) {
    return 1, 8, "slashing"
}

func (l *Longsword) Action() string {
    return "versatile"
}

func (l *Longsword) GetQuantity() int {
    return 1
}

func (l *Longsword) GetRange() []float64 {
    return []float64{0, 0}
}
