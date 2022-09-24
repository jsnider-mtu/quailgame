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

func (d Darts) Use() {
    // must be equipped to use
}

func (d Darts) Save() string {
    return "Darts," + strconv.Itoa(d.Quantity)
}

func (d Darts) PrettyPrint() string {
    return fmt.Sprintf("Darts%17s", "Quantity: " + strconv.Itoa(d.Quantity))
}
