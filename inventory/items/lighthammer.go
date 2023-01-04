package items

import (
    "fmt"
    "strconv"
)

type LightHammer struct {
    Quantity int
}

func (l *LightHammer) Slot() string {
    return "RightHand"
}

func (l *LightHammer) Use() (string, []int) {
    return l.Action(), []int{96, 288}
}

func (l *LightHammer) Save() string {
    return "LightHammer," + strconv.Itoa(l.Quantity)
}

func (l *LightHammer) PrettyPrint() string {
    return fmt.Sprintf("Light Hammer (%d)", l.Quantity)
}

func (l *LightHammer) Function() string {
    return "melee-light-throw"
}

func (l *LightHammer) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (l *LightHammer) Action() string {
    return "throw"
}

func (l *LightHammer) GetQuantity() int {
    return l.Quantity
}

func (l *LightHammer) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
