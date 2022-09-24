package items

type Flail struct {
}

func (f Flail) Slot() string {
    return "RightHand"
}

func (f Flail) Use() {
    // must be equipped to use
}

func (f Flail) Save() string {
    return "Flail"
}

func (f Flail) PrettyPrint() string {
    return "Flail"
}
