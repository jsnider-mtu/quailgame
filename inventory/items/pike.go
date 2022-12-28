package items

type Pike struct {
}

func (p Pike) Slot() string {
    return "BothHands"
}

func (p Pike) Use() {
    // must be equipped to use
}

func (p Pike) Save() string {
    return "Pike"
}

func (p Pike) PrettyPrint() string {
    return "Pike"
}

func (p Pike) Function() string {
    return "melee"
}

func (p Pike) Damage() (int, int, string) {
    return 1, 10, "piercing"
}
