package items

type Scimitar struct {
}

func (s *Scimitar) Slot() string {
    return "RightHand"
}

func (s *Scimitar) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (s *Scimitar) Save() string {
    return "Scimitar"
}

func (s *Scimitar) PrettyPrint() string {
    return "Scimitar"
}

func (s *Scimitar) Function() string {
    return "melee"
}

func (s *Scimitar) Damage() (int, int, string) {
    return 1, 6, "slashing"
}

func (s *Scimitar) Action() string {
    return ""
}

func (s *Scimitar) GetQuantity() int {
    return 1
}

func (s *Scimitar) GetRange() []float64 {
    return []float64{0, 0}
}
