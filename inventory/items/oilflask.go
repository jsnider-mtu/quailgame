package items

import "strconv"

type Oilflask struct {
    quantity int
}

func (o Oilflask) Slot() string {
    return ""
}

func (o Oilflask) Use() {
}

func (o Oilflask) Save() string {
    return "Oilflask," + strconv.Itoa(o.quantity)
}
