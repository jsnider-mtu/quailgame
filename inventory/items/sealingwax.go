package items

type SealingWax struct {
}

func (s *SealingWax) Slot() string {
    return ""
}

func (s *SealingWax) Use() (string, []int) {
    return "", []int{}
}

func (s *SealingWax) Save() string {
    return "SealingWax"
}

func (s *SealingWax) PrettyPrint() string {
    return "Sealing Wax"
}

func (s *SealingWax) Function() string {
    return "writing"
}

func (s *SealingWax) Damage() (int, int, string) {
    return 0, 0, ""
}

func (s *SealingWax) Action() string {
    return ""
}

func (s *SealingWax) GetQuantity() int {
    return 1
}
