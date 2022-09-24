package items

import "strconv"

type Torches struct {
    Quantity int
}

func (t Torches) Slot() string {
    return ""
}

func (t Torches) Use() {
}

func (t Torches) Save() string {
    return "Torches," + strconv.Itoa(t.Quantity)
}

func (t Torches) PrettyPrint() string {
    return "Torches\t\tQuantity: " + strconv.Itoa(t.Quantity)
}
