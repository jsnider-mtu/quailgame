package items

import "strconv"

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
