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
    return "OilFlask," + strconv.Itoa(o.Quantity)
}

func (o OilFlask) PrettyPrint() string {
    return fmt.Sprintf("Oil Flask (%d)", o.Quantity)
}
