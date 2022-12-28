package items

type Blowgun struct {
}

func (b Blowgun) Slot() string {
    return "RightHand"
}

func (b Blowgun) Use() (string, []int) {
    // must be equipped to use
}

func (b Blowgun) Save() string {
    return "Blowgun"
}

func (b Blowgun) PrettyPrint() string {
    return "Blowgun"
}

func (b Blowgun) Function() string {
    return "range"
}

func (b Blowgun) Damage() (int, int, string) {
    return 1, 1, "piercing"
}
