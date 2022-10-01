package items

type Panflute struct {
}

func (p PanFlute) Slot() string {
    return "BothHands"
}

func (p PanFlute) Use() {
}

func (p PanFlute) Save() string {
    return "PanFlute"
}

func (p PanFlute) PrettyPrint() string {
    return "Pan Flute"
}
