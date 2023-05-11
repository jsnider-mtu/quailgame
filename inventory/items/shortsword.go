package items

type Shortsword struct {
}

func (s *Shortsword) Slot() string {
    return "RightHand"
}

func (s *Shortsword) Use() (string, []int) {
    return s.Action(), []int{}
}

func (s *Shortsword) Save() string {
    return "Shortsword"
}

func (s *Shortsword) PrettyPrint() string {
    return "Shortsword"
}

func (s *Shortsword) Function() string {
    return "melee-finesse-light"
}

func (s *Shortsword) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (s *Shortsword) Action() string {
    return ""
}

func (s *Shortsword) GetQuantity() int {
    return 1
}

func (s *Shortsword) GetRange() []float64 {
    return []float64{0, 0}
}
