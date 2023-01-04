package items

type Handaxe struct {
}

func (h *Handaxe) Slot() string {
    return "RightHand"
}

func (h *Handaxe) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (h *Handaxe) Save() string {
    return "Handaxe"
}

func (h *Handaxe) PrettyPrint() string {
    return "Handaxe"
}

func (h *Handaxe) Function() string {
    return "melee"
}

func (h *Handaxe) Damage() (int, int, string) {
    return 1, 6, "slashing"
}

func (h *Handaxe) Action() string {
    return ""
}

func (h *Handaxe) GetQuantity() int {
    return 1
}

func (h *Handaxe) GetRange() []float64 {
    return []float64{0, 0}
}
