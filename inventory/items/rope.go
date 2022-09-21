package items

import "strconv"

type Rope struct {
    length int
}

func (r Rope) Slot() string {
    return ""
}

func (r Rope) Use() {
}

func (r Rope) Save() string {
    return "Rope," + strconv.Itoa(r.length)
}
