package items

import "strconv"

type Paper struct {
    Quantity int
}

func (p Paper) Slot() string {
    return ""
}

func (p Paper) Use() {
}

func (p Paper) Save() string {
    return "Paper," + strconv.Itoa(p.Quantity)
}

func (p Paper) PrettyPrint() string {
    return "Paper\t\tQuantity: " + strconv.Itoa(p.Quantity)
}
