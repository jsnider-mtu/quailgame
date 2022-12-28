package items

import (
    "fmt"
    "strconv"
)

type Darts struct {
    Quantity int
}

func (d Darts) Slot() string {
    return "RightHand"
}

func (d Darts) Use() (string, []int) {
    // must be equipped to use
}

func (d Darts) Save() string {
    return "Darts," + strconv.Itoa(d.Quantity)
}

func (d Darts) PrettyPrint() string {
    return fmt.Sprintf("Darts (%d)", d.Quantity)
}

func (d Darts) Function() string {
    return "range"
}

func (d Darts) Damage() (int, int, string) {
    return 1, 4, "piercing"
}
