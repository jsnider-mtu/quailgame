package items

type Lyre struct {
}

func (l Lyre) Slot() string {
    return "BothHands"
}

func (l Lyre) Use() {
}

func (l Lyre) Save() string {
    return "Lyre"
}

func (l Lyre) PrettyPrint() string {
    return "Lyre"
}

func (l Lyre) Function() string {
    return "instrument"
}

func (l Lyre) Damage() (int, int, string) {
    return 0, 0, ""
}
