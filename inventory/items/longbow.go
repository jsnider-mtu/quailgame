package items

import (
    "fmt"
    "strconv"
)

type Longbow struct {
    Ammo int
}

func (l *Longbow) Slot() string {
    return "BothHands"
}

func (l *Longbow) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *Longbow) Save() string {
    return "Longbow," + strconv.Itoa(l.Ammo)
}

func (l *Longbow) PrettyPrint() string {
    return fmt.Sprintf("Longbow (%d)", l.Ammo)
}

func (l *Longbow) Function() string {
    return "range-ammo-heavy"
}

func (l *Longbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (l *Longbow) Action() string {
    return ""
}

func (l *Longbow) GetQuantity() int {
    return l.Ammo
}

func (l *Longbow) GetRange() []float64 {
    return []float64{720.0, 2880.0}
}
