package items

import (
    "fmt"
    "strconv"
)

type OilFlask struct {
    Quantity int
}

func (o OilFlask) Slot() string {
    return "LeftHand"
}

func (o OilFlask) Use() {
}

func (o OilFlask) Save() string {
    return "Oil Flask," + strconv.Itoa(o.Quantity)
}

func (o OilFlask) PrettyPrint() string {
    return fmt.Sprintf("Oil Flask%14s", "Quantity: " + strconv.Itoa(o.Quantity))
}
