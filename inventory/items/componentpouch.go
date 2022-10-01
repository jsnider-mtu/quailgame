package items

type ComponentPouch struct {
}

func (c ComponentPouch) Slot() string {
    return "LeftHand"
}

func (c ComponentPouch) Use() {
}

func (c ComponentPouch) Save() string {
    return "ComponentPouch"
}

func (c ComponentPouch) PrettyPrint() string {
    return "Component Pouch"
}
