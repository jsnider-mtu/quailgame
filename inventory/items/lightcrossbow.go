package items

import (
    "fmt"
    "strconv"
)

type LightCrossbow struct {
    Ammo int
}

func (l *LightCrossbow) Slot() string {
    return "BothHands"
}

func (l *LightCrossbow) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *LightCrossbow) Save() string {
    return "LightCrossbow," + strconv.Itoa(l.Ammo)
}

func (l *LightCrossbow) PrettyPrint() string {
    return fmt.Sprintf("Light Crossbow (%d)", l.Ammo)
}

func (l *LightCrossbow) Function() string {
    return "range-ammo-loading"
}

func (l *LightCrossbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}

func (l *LightCrossbow) Action() string {
    return ""
}

func (l *LightCrossbow) GetQuantity() int {
    return l.Quantity
}

func (l *LightCrossbow) GetRange() []float64 {
    return []float64{384.0, 1536.0}
}
