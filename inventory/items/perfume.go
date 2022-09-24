package items

type Perfume struct {
}

func (p Perfume) Slot() string {
    return ""
}

func (p Perfume) Use() {
}

func (p Perfume) Save() string {
    return "Perfume"
}

func (p Perfume) PrettyPrint() string {
    return "Perfume"
}
