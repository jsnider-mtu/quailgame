package items

type Lance struct {
}

func (l *Lance) Slot() string {
    return "RightHand"
}

func (l *Lance) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (l *Lance) Save() string {
    return "Lance"
}

func (l *Lance) PrettyPrint() string {
    return "Lance"
}

func (l *Lance) Function() string {
    return "melee"
}

func (l *Lance) Damage() (int, int, string) {
    return 1, 12, "piercing"
}

func (l *Lance) Action() string {
    return ""
}

func (l *Lance) GetQuantity() int {
    return 1
}
