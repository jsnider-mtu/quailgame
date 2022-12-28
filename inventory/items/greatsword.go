package items

type Greatsword struct {
}

func (g Greatsword) Slot() string {
    return "BothHands"
}

func (g Greatsword) Use() (string, []int) {
    // must be equipped to use
}

func (g Greatsword) Save() string {
    return "Greatsword"
}

func (g Greatsword) PrettyPrint() string {
    return "Greatsword"
}

func (g Greatsword) Function() string {
    return "melee"
}

func (g Greatsword) Damage() (int, int, string) {
    return 2, 6, "slashing"
}
