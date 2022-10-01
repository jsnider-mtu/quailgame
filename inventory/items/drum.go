package items

type Drum struct {
}

func (d Drum) Slot() string {
    return "BothHands"
}

func (d Drum) Use() {
}

func (d Drum) Save() string {
    return "Drum"
}

func (d Drum) PrettyPrint() string {
    return "Drum"
}
