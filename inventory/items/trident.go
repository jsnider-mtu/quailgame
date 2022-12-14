package items

type Trident struct {
}

func (t *Trident) Slot() string {
    return "BothHands"
}

func (t *Trident) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (t *Trident) Save() string {
    return "Trident"
}

func (t *Trident) PrettyPrint() string {
    return "Trident"
}

func (t *Trident) Function() string {
    return "melee"
}

func (t *Trident) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (t *Trident) Action() string {
    return ""
}

func (t *Trident) GetQuantity() int {
    return 1
}

func (t *Trident) GetRange() []float64 {
    return []float64{0, 0}
}
