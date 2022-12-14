package items

type Spear struct {
}

func (s *Spear) Slot() string {
    return "BothHands"
}

func (s *Spear) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (s *Spear) Save() string {
    return "Spear"
}

func (s *Spear) PrettyPrint() string {
    return "Spear"
}

func (s *Spear) Function() string {
    return "melee"
}

func (s *Spear) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (s *Spear) Action() string {
    return ""
}

func (s *Spear) GetQuantity() int {
    return 1
}

func (s *Spear) GetRange() []float64 {
    return []float64{0, 0}
}
