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

func (d Drum) Function() string {
    return "instrument"
}

func (d Drum) Damage() (int, int, string) {
    return 0, 0, ""
}
