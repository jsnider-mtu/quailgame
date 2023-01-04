package items

type Sling struct {
}

func (s *Sling) Slot() string {
    return "RightHand"
}

func (s *Sling) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (s *Sling) Save() string {
    return "Sling"
}

func (s *Sling) PrettyPrint() string {
    return "Sling"
}

func (s *Sling) Function() string {
    return "range"
}

func (s *Sling) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (s *Sling) Action() string {
    return ""
}

func (s *Sling) GetQuantity() int {
    return 1
}

func (s *Sling) GetRange() []float64 {
    return []float64{0, 0}
}
