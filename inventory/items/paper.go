package items

type Paper struct {
    quantity int
}

func (p Paper) Slot() string {
    return ""
}

func (p Paper) Use() {
}

func (p Paper) Save() string {
    return "Paper"
}
