package items

type Shortbow struct {
}

func (s *Shortbow) Slot() string {
    return "BothHands"
}

func (s *Shortbow) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (s *Shortbow) Save() string {
    return "Shortbow"
}

func (s *Shortbow) PrettyPrint() string {
    return "Shortbow"
}

func (s *Shortbow) Function() string {
    return "range"
}

func (s *Shortbow) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (s *Shortbow) Action() string {
    return ""
}

func (s *Shortbow) GetQuantity() int {
    return 1
}

func (s *Shortbow) GetRange() []float64 {
    return []float64{0, 0}
}
