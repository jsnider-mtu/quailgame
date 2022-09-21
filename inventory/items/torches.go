package items

import "strconv"

type Torches struct {
    quantity int
}

func (t Torches) Slot() string {
    return ""
}

func (t Torches) Use() {
}

func (t Torches) Save() string {
    return "Torches," + strconv.Itoa(t.quantity)
}
