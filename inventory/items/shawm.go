package items

type Shawm struct {
}

func (s *Shawm) Slot() string {
    return "BothHands"
}

func (s *Shawm) Use() (string, []int) {
    return "", []int{}
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
    return ""
}
