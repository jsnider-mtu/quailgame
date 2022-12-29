package items

import (
    "fmt"
    "strconv"
)

type OilFlask struct {
    Quantity int
}

func (o *OilFlask) Slot() string {
    return "LeftHand"
}

func (o *OilFlask) Use() (string, []int) {
    return "", []int{}
}

func (o *OilFlask) Save() string {
    return "OilFlask," + strconv.Itoa(o.Quantity)
}

func (o *OilFlask) PrettyPrint() string {
    return fmt.Sprintf("Oil Flask (%d)", o.Quantity)
}

func (o *OilFlask) Function() string {
    return "fire"
}

func (o *OilFlask) Damage() (int, int, string) {
    return 0, 0, ""
}

func (o *OilFlask) Action() string {
    return ""
}
