package items

import (
    "fmt"
    "strconv"
)

type OilFlask struct {
    Quantity int
}

func (o *OilFlask) Slot() string {
    return "RightHand"
}

func (o *OilFlask) Use() (string, []int) {
    return o.Action(), []int{96, 96}
}

func (o *OilFlask) Save() string {
    return "OilFlask," + strconv.Itoa(o.Quantity)
}

func (o *OilFlask) PrettyPrint() string {
    return fmt.Sprintf("Oil Flask (%d)", o.Quantity)
}

func (o *OilFlask) Function() string {
    return "range"
}

func (o *OilFlask) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (o *OilFlask) Action() string {
    if o.Quantity > 0 {
        return "throw"
    } else {
        return ""
    }
}

func (o *OilFlask) GetQuantity() int {
    return o.Quantity
}

func (o *OilFlask) GetRange() []float64 {
    return []float64{96.0, 96.0}
}
