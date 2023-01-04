package items

import (
    "fmt"
    "strconv"
)

type Lamp struct {
    Quantity int
    Turns int
}

func (l *Lamp) Slot() string {
    return "LeftHand"
}

func (l *Lamp) Use() (string, []int) {
    return l.Action(), []int{15, 30, l.Turns}
}

func (l *Lamp) Save() string {
    return "Lamp," + strconv.Itoa(l.Turns) + "," + strconv.Itoa(l.Quantity)
}

func (l *Lamp) PrettyPrint() string {
    return fmt.Sprintf("Lamp (%d)", l.Quantity)
}

func (l *Lamp) Function() string {
    return "light"
}

func (l *Lamp) Damage() (int, int, string) {
    return 0, 0, ""
}

func (l *Lamp) Action() string {
    if l.Quantity > 0 {
        return "illuminate"
    }
    return ""
}

func (l *Lamp) GetQuantity() int {
    return l.Quantity
}

func (l *Lamp) GetRange() []float64 {
    return []float64{0, 0}
}
