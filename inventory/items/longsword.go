package items

type Longsword struct {
}

func (l Longsword) Slot() string {
    return "RightHand"
}

func (l Longsword) Use() (string, []int) {
    // must be equipped to use
}

func (l Longsword) Save() string {
    return "Longsword"
}

func (l Longsword) PrettyPrint() string {
    return "Longsword"
}

func (l Longsword) Function() string {
    return "melee"
}

func (l Longsword) Damage() (int, int, string) {
    return 1, 8, "slashing"
}
