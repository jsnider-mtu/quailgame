package items

type Sickle struct {
}

func (s *Sickle) Slot() string {
    return "RightHand"
}

func (s *Sickle) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (s *Sickle) Save() string {
    return "Sickle"
}

func (s *Sickle) PrettyPrint() string {
    return "Sickle"
}

func (s *Sickle) Function() string {
    return "melee"
}

func (s *Sickle) Damage() (int, int, string) {
    return 1, 4, "slashing"
}

func (s *Sickle) Action() string {
    return ""
}

func (s *Sickle) GetQuantity() int {
    return 1
}

func (s *Sickle) GetRange() []float64 {
    return []float64{0, 0}
}
