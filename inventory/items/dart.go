package items

import "strconv"

type Dart struct {
    Quantity int
}

func (d Dart) Slot() string {
    return "RightHand"
}

func (d Dart) Use() {
    // must be equipped to use
}

func (d Dart) Save() string {
    return "Dart," + strconv.Itoa(d.Quantity)
}
