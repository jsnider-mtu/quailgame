package items

type Shield struct {
}

func (s Shield) Slot() string {
    return "LeftHand"
}

func (s Shield) Use() (string, []int) {
}

func (s Shield) Save() string {
    return "Shield"
}

func (s Shield) PrettyPrint() string {
    return "Shield"
}

func (s Shield) Function() string {
    return "armor"
}

func (s Shield) Damage() (int, int, string) {
    return 0, 0, ""
}
