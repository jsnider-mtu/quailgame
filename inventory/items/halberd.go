package items

type Halberd struct {
}

func (h *Halberd) Slot() string {
    return "BothHands"
}

func (h *Halberd) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (h *Halberd) Save() string {
    return "Halberd"
}

func (h *Halberd) PrettyPrint() string {
    return "Halberd"
}

func (h *Halberd) Function() string {
    return "melee"
}

func (h *Halberd) Damage() (int, int, string) {
    return 1, 10, "slashing"
}

func (h *Halberd) Action() string {
    return ""
}

func (h *Halberd) GetQuantity() int {
    return 1
}

func (h *Halberd) GetRange() []float64 {
    return []float64{0, 0}
}
