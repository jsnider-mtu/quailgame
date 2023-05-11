package items

import (
    "fmt"
    "strconv"
)

type Sling struct {
    Ammo int
}

func (s *Sling) Slot() string {
    return "RightHand"
}

func (s *Sling) Use() (string, []int) {
    return s.Action(), []int{}
}

func (s *Sling) Save() string {
    return "Sling," + strconv.Itoa(s.Ammo)
}

func (s *Sling) PrettyPrint() string {
    return fmt.Sprintf("Sling (%d)", s.Ammo)
}

func (s *Sling) Function() string {
    return "range-ammo"
}

func (s *Sling) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (s *Sling) Action() string {
    return ""
}

func (s *Sling) GetQuantity() int {
    return s.Ammo
}

func (s *Sling) GetRange() []float64 {
    return []float64{144.0, 576.0}
}
