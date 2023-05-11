package items

type Blowgun struct {
    Ammo int
}

func (b *Blowgun) Slot() string {
    return "RightHand"
}

func (b *Blowgun) Use() (string, []int) {
    return b.Action(), []int{}
}

func (b *Blowgun) Save() string {
    return "Blowgun"
}

func (b *Blowgun) PrettyPrint() string {
    return "Blowgun"
}

func (b *Blowgun) Function() string {
    return "range-ammo-loading"
}

func (b *Blowgun) Damage() (int, int, string) {
    return 1, 1, "piercing"
}

func (b *Blowgun) Action() string {
    return ""
}

func (b *Blowgun) GetQuantity() int {
    return b.Ammo
}

func (b *Blowgun) GetRange() []float64 {
    return []float64{120.0, 480.0}
}
