package items

type Shawm struct {
}

func (s *Shawm) Slot() string {
    return "BothHands"
}

func (s *Shawm) Use() (string, []int) {
    return s.Action(), []int{}
}

func (s *Shawm) Save() string {
    return "Shawm"
}

func (s *Shawm) PrettyPrint() string {
    return "Shawm"
}

func (s *Shawm) Function() string {
    return "instrument"
}

func (s *Shawm) Damage() (int, int, string) {
    return 0, 0, ""
}

func (s *Shawm) Action() string {
    return "playmusic"
}

func (s *Shawm) GetQuantity() int {
    return 1
}

func (s *Shawm) GetRange() []float64 {
    return []float64{0, 0}
}
