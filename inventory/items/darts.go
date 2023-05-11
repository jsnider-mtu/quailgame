package items

import (
    "fmt"
    "strconv"
)

type Darts struct {
    Quantity int
}

func (d *Darts) Slot() string {
    return "RightHand"
}

func (d *Darts) Use() (string, []int) {
    return d.Action(), []int{96, 288}
}

func (d *Darts) Save() string {
    return "Darts," + strconv.Itoa(d.Quantity)
}

func (d *Darts) PrettyPrint() string {
    return fmt.Sprintf("Darts (%d)", d.Quantity)
}

func (d *Darts) Function() string {
    return "range-finesse-throw"
}

func (d *Darts) Damage() (int, int, string) {
    return 1, 4, "piercing"
}

func (d *Darts) Action() string {
    return "throw"
}

func (d *Darts) GetQuantity() int {
    return d.Quantity
}

func (d *Darts) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
