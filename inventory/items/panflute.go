package items

type PanFlute struct {
}

func (p PanFlute) Slot() string {
    return "BothHands"
}

func (p PanFlute) Use() (string, []int) {
    return "", []int{}
}

func (p PanFlute) Save() string {
    return "PanFlute"
}

func (p PanFlute) PrettyPrint() string {
    return "Pan Flute"
}

func (p PanFlute) Function() string {
    return "instrument"
}

func (p PanFlute) Damage() (int, int, string) {
    return 0, 0, ""
}
