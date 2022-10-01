package items

import (
    "fmt"
    "strconv"
)

type Oilflask struct {
    Quantity int
}

func (o Oilflask) Slot() string {
    return ""
}

func (o Oilflask) Use() {
}

func (o Oilflask) Save() string {
    return "Oilflask," + strconv.Itoa(o.Quantity)
}

func (o Oilflask) PrettyPrint() string {
    return fmt.Sprintf("Oilflask%14s", "Quantity: " + strconv.Itoa(o.Quantity))
}
