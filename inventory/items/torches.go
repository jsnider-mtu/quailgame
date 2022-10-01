package items

import (
    "fmt"
    "strconv"
)

type Torches struct {
    Quantity int
}

func (t Torches) Slot() string {
    return "LeftHand"
}

func (t Torches) Use() {
}

func (t Torches) Save() string {
    return "Torches," + strconv.Itoa(t.Quantity)
}

func (t Torches) PrettyPrint() string {
    return fmt.Sprintf("Torches%15s", "Quantity: " + strconv.Itoa(t.Quantity))
}
