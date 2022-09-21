package items

import "strconv"

type Darts struct {
    Quantity int
}

func (d Darts) Slot() string {
    return "RightHand"
}

func (d Darts) Use() {
    // must be equipped to use
}

func (d Darts) Save() string {
    return "Dart," + strconv.Itoa(d.Quantity)
}
