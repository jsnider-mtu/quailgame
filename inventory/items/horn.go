package items

type Horn struct {
}

func (h *Horn) Slot() string {
    return "BothHands"
}

func (h *Horn) Use() (string, []int) {
    return "", []int{}
}

func (h *Horn) Save() string {
    return "Horn"
}

func (h *Horn) PrettyPrint() string {
    return "Horn"
}

func (h *Horn) Function() string {
    return "instrument"
}

func (h *Horn) Damage() (int, int, string) {
    return 0, 0, ""
}

func (h *Horn) Action() string {
    return ""
}

func (h *Horn) GetQuantity() int {
    return 1
}

func (h *Horn) GetRange() []float64 {
    return []float64{0, 0}
}
