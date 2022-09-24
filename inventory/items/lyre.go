package items

type Lyre struct {
}

func (l Lyre) Slot() string {
    return ""
}

func (l Lyre) Use() {
}

func (l Lyre) Save() string {
    return "Lyre"
}

func (l Lyre) PrettyPrint() string {
    return "Lyre"
}
