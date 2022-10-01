package items

type ThievesTools struct {
}

func (t ThievesTools) Slot() string {
    return ""
}

func (t ThievesTools) Use() {
}

func (t ThievesTools) Save() string {
    return "ThievesTools"
}

func (t ThievesTools) PrettyPrint() string {
    return "Thieves Tools"
}
