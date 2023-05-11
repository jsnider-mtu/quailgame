package items

type Shield struct {
}

func (s *Shield) Slot() string {
    return "LeftHand"
}

func (s *Shield) Use() (string, []int) {
    return s.Action(), []int{48}
}

func (s *Shield) Save() string {
    return "Shield"
}

func (s *Shield) PrettyPrint() string {
    return "Shield"
}

func (s *Shield) Function() string {
    return "armor"
}

func (s *Shield) Damage() (int, int, string) {
    return 0, 0, ""
}

func (s *Shield) Action() string {
    return "bash"
}

func (s *Shield) GetQuantity() int {
    return 1
}

func (s *Shield) GetRange() []float64 {
    return []float64{0, 0}
}
