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

func (t ThievesTools) Function() string {
    return "theft"
}

func (t ThievesTools) Damage() (int, int, string) {
    return 0, 0, ""
}
