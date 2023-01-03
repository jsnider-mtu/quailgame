package items

import (
    "fmt"
    "strconv"
)

type Lamp struct {
    Quantity int
}

func (l *Lamp) Slot() string {
    return "LeftHand"
}

func (l *Lamp) Use() (string, []int) {
    if l.Quantity > 0 {
        return "illuminate", []int{15, 30, 3600}
    }
    return "", []int{}
}

func (l *Lamp) Save() string {
    return "Lamp," + strconv.Itoa(l.Quantity)
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
    return "illuminate"
}

func (l *Lamp) GetQuantity() int {
    return l.Quantity
}
