package items

import (
    "fmt"
    "strconv"
)

type Candles struct {
    Quantity int
}

func (c Candles) Slot() string {
    return "LeftHand"
}

func (c Candles) Use() {
}

func (c Candles) Save() string {
    return "Candles," + strconv.Itoa(c.Quantity)
}

func (c Candles) PrettyPrint() string {
    return fmt.Sprintf("Candles%15s", "Quantity: " + strconv.Itoa(c.Quantity))
}
