package items

type ArcaneFocus struct {
}

func (a *ArcaneFocus) Slot() string {
    return "LeftHand"
}

func (a *ArcaneFocus) Use() (string, []int) {
    return "", []int{}
}

func (a *ArcaneFocus) Save() string {
    return "ArcaneFocus"
}

func (a *ArcaneFocus) PrettyPrint() string {
    return "Arcane Focus" 
}

func (a *ArcaneFocus) Function() string {
    return "spells"
}

func (a *ArcaneFocus) Damage() (int, int, string) {
    return 0, 0, ""
}

func (a *ArcaneFocus) Action() string {
    return ""
}

func (a *ArcaneFocus) GetQuantity() int {
    return 1
}

func (a *ArcaneFocus) GetRange() []float64 {
    return []float64{0, 0}
}
