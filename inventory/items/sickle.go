package items

type Sickle struct {
}

func (s Sickle) Slot() string {
    return "RightHand"
}

func (s Sickle) Use() (string, []int) {
    // must be equipped to use
}

func (s Sickle) Save() string {
    return "Sickle"
}

func (s Sickle) PrettyPrint() string {
    return "Sickle"
}

func (s Sickle) Function() string {
    return "melee"
}

func (s Sickle) Damage() (int, int, string) {
    return 1, 4, "slashing"
}
