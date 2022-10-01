package items

type Horn struct {
}

func (h Horn) Slot() string {
    return ""
}

func (h Horn) Use() {
}

func (h Horn) Save() string {
    return "Horn"
}

func (h Horn) PrettyPrint() string {
    return "Horn"
}
