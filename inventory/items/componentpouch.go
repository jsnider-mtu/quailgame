package items

type ComponentPouch struct {
}

func (c ComponentPouch) Slot() string {
    return "LeftHand"
}

func (c ComponentPouch) Use() (string, []int) {
}

func (c ComponentPouch) Save() string {
    return "ComponentPouch"
}

func (c ComponentPouch) PrettyPrint() string {
    return "Component Pouch"
}

func (c ComponentPouch) Function() string {
    return "spells"
}

func (c ComponentPouch) Damage() (int, int, string) {
    return 0, 0, ""
}
