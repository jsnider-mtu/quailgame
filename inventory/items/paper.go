package items

import "strconv"

type Paper struct {
    quantity int
}

func (p Paper) Slot() string {
    return ""
}

func (p Paper) Use() {
}

func (p Paper) Save() string {
    return "Paper," + strconv.Itoa(p.quantity)
}
