package items

type Scalemail struct {
}

func (s *Scalemail) Slot() string {
    return "Armor"
}

func (s *Scalemail) Use() (string, []int) {
    return "", []int{}
}

func (s *Scalemail) Save() string {
    return "Scalemail"
}

func (s *Scalemail) PrettyPrint() string {
    return "Scalemail"
}

func (s *Scalemail) Function() string {
    return "armor"
}

func (s *Scalemail) Damage() (int, int, string) {
    return 0, 0, ""
}

func (s *Scalemail) Action() string {
    return ""
}

func (s *Scalemail) GetQuantity() int {
    return 1
}

func (s *Scalemail) GetRange() []float64 {
    return []float64{0, 0}
}
