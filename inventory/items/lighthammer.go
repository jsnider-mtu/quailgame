package items

type Lighthammer struct {
}

func (l Lighthammer) Slot() string {
    return "RightHand"
}

func (l Lighthammer) Use() {
    // must be equipped to use
}

func (l Lighthammer) Save() string {
    return "Lighthammer"
}
