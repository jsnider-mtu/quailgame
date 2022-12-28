package items

type Dagger struct {
}

func (d Dagger) Slot() string {
    return "RightHand"
}

func (d Dagger) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (d Dagger) Save() string {
    return "Dagger"
}

func (d Dagger) PrettyPrint() string {
    return "Dagger"
}

func (d Dagger) Function() string {
    return "melee"
}

func (d Dagger) Damage() (int, int, string) {
    return 1, 4, "piercing"
}
