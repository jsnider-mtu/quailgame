package items

import (
    "fmt"
    "strconv"
)

type Torches struct {
    Quantity int
}

func (t *Torches) Slot() string {
    return "LeftHand"
}

func (t *Torches) Use() (string, []int) {
    return "", []int{}
}

func (t *Torches) Save() string {
    return "Torches," + strconv.Itoa(t.Quantity)
}

func (t *Torches) PrettyPrint() string {
    return fmt.Sprintf("Torches (%d)", t.Quantity)
}

func (t *Torches) Function() string {
    return "light"
}

func (t *Torches) Damage() (int, int, string) {
    return 0, 0, ""
}

func (t *Torches) Action() string {
    return ""
}

func (t *Torches) GetQuantity() int {
    return t.Quantity
}

func (t *Torches) GetRange() []float64 {
    return []float64{0, 0}
}
