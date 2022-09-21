package items

type Lightcrossbow struct {
}

func (l Lightcrossbow) Slot() string {
    return "RightHand"
}

func (l Lightcrossbow) Use() {
    // must be equipped to use
}

func (l Lightcrossbow) Save() string {
    return "Lightcrossbow"
}
