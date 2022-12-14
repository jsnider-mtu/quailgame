package items

import (
    "fmt"
    "strconv"
)

type Candles struct {
    Quantity int
    Turns int
}

func (c *Candles) Slot() string {
    return "LeftHand"
}

func (c *Candles) Use() (string, []int) {
    return c.Action(), []int{5, 5, c.Turns}
}

func (c *Candles) Save() string {
    return "Candles," + strconv.Itoa(c.Turns) + "," + strconv.Itoa(c.Quantity)
}

func (c *Candles) PrettyPrint() string {
    return fmt.Sprintf("Candles (%d)", c.Quantity)
}

func (c *Candles) Function() string {
    return "light"
}

func (c *Candles) Damage() (int, int, string) {
    return 0, 0, ""
}

func (c *Candles) Action() string {
    if c.Quantity > 0 {
        return "illuminate"
    }
    return ""
}

func (c *Candles) GetQuantity() int {
    return c.Quantity
}

func (c *Candles) GetRange() []float64 {
    return []float64{0, 0}
}
