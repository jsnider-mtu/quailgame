package items

type Leatherarmor struct {
}

func (l Leatherarmor) Slot() string {
    return ""
}

func (l Leatherarmor) Use() {
}

func (l Leatherarmor) Save() string {
    return "Leatherarmor"
}

func (l Leatherarmor) PrettyPrint() string {
    return "Leatherarmor"
}
