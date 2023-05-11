package items

type Battleaxe struct {
    Currslot string
}

func (b *Battleaxe) GetCurrSlot() string {
    return b.Currslot
}

func (b *Battleaxe) SwitchSlots() {
    if b.Currslot == "BothHands" {
        b.Currslot = "RightHand"
    } else if b.Currslot == "RightHand" {
        b.Currslot = "BothHands"
    }
    return
}

func (b *Battleaxe) Slot() string {
    return "Versatile"
}

func (b *Battleaxe) Use() (string, []int) {
    return b.Action(), []int{1, 10}
}

func (b *Battleaxe) Save() string {
    return "Battleaxe"
}

func (b *Battleaxe) PrettyPrint() string {
    return "Battleaxe"
}

func (b *Battleaxe) Function() string {
    return "melee"
}

func (b *Battleaxe) Damage() (int, int, string) {
    return 1, 8, "slashing"
}

func (b *Battleaxe) Action() string {
    return "versatile"
}

func (b *Battleaxe) GetQuantity() int {
    return 1
}

func (b *Battleaxe) GetRange() []float64 {
    return []float64{0, 0}
}
