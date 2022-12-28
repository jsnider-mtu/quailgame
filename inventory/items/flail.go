package items

type Flail struct {
}

func (f Flail) Slot() string {
    return "RightHand"
}

func (f Flail) Use() (string, []int) {
    // must be equipped to use
}

func (f Flail) Save() string {
    return "Flail"
}

func (f Flail) PrettyPrint() string {
    return "Flail"
}

func (f Flail) Function() string {
    return "melee"
}

func (f Flail) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}
