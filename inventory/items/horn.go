package items

type Horn struct {
}

func (h Horn) Slot() string {
    return "BothHands"
}

func (h Horn) Use() {
}

func (h Horn) Save() string {
    return "Horn"
}

func (h Horn) PrettyPrint() string {
    return "Horn"
}

func (h Horn) Function() string {
    return "instrument"
}

func (h Horn) Damage() (int, int, string) {
    return 0, 0, ""
}
