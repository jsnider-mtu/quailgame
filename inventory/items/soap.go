package items

type Soap struct {
}

func (s *Soap) Slot() string {
    return ""
}

func (s *Soap) Use() (string, []int) {
    return "", []int{}
}

func (s *Soap) Save() string {
    return "Soap"
}

func (s *Soap) PrettyPrint() string {
    return "Soap"
}

func (s *Soap) Function() string {
    return "cleaning"
}

func (s *Soap) Damage() (int, int, string) {
    return 0, 0, ""
}

func (s *Soap) Action() string {
    return ""
}
