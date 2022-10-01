package items

type Soap struct {
}

func (s Soap) Slot() string {
    return ""
}

func (s Soap) Use() {
}

func (s Soap) Save() string {
    return "Soap"
}

func (s Soap) PrettyPrint() string {
    return "Soap"
}
