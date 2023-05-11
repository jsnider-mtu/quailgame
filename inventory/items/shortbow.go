package items

import (
    "fmt"
    "strconv"
)

type Shortbow struct {
    Ammo int
}

func (s *Shortbow) Slot() string {
    return "BothHands"
}

func (s *Shortbow) Use() (string, []int) {
    return s.Action(), []int{}
}

func (s *Shortbow) Save() string {
    return "Shortbow," + strconv.Itoa(s.Ammo)
}

func (s *Shortbow) PrettyPrint() string {
    return fmt.Sprintf("Shortbow (%d)", s.Ammo)
}

func (s *Shortbow) Function() string {
    return "range-ammo"
}

func (s *Shortbow) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (s *Shortbow) Action() string {
    return ""
}

func (s *Shortbow) GetQuantity() int {
    return s.Ammo
}

func (s *Shortbow) GetRange() []float64 {
    return []float64{384.0, 1536.0}
}
