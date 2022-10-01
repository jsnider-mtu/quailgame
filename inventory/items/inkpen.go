package items

type Inkpen struct {
}

func (i Inkpen) Slot() string {
    return ""
}

func (i Inkpen) Use() {
}

func (i Inkpen) Save() string {
    return "Inkpen"
}

func (i Inkpen) PrettyPrint() string {
    return "Inkpen"
}
