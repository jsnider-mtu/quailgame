package items

type Battleaxe struct {
}

func (b *Battleaxe) Slot() string {
    return "BothHands"
}

func (b *Battleaxe) Use() (string, []int) {
    return b.Action(), []int{}
}

func (b *Battleaxe) Save() string {
    return "Battleaxe"
}

func (b *Battleaxe) PrettyPrint() string {
    return "Battleaxe"
}

func (b *Battleaxe) Function() string {
    return "melee-versatile"
}

func (b *Battleaxe) Damage() (int, int, string) {
    return 1, 8, "slashing"
}

func (b *Battleaxe) Action() string {
    return ""
}

func (b *Battleaxe) GetQuantity() int {
    return 1
}

func (b *Battleaxe) GetRange() []float64 {
    return []float64{0, 0}
}
