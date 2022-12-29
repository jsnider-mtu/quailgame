package items

import (
    "fmt"
    "log"
    "strconv"
)

type Candles struct {
    Quantity int
}

func (c *Candles) Slot() string {
    return "LeftHand"
}

func (c *Candles) Use() (string, []int) {
    // illuminate surroundings
    if c.Quantity > 0 {
        c.Quantity = c.Quantity - 1
        log.Println(fmt.Sprintf("c.Quantity == %d", c.Quantity))
        return "illuminate", []int{5, 5, 600}
    }
    return "", []int{}
}

func (c *Candles) Save() string {
    return "Candles," + strconv.Itoa(c.Quantity)
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
