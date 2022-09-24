package items

type Scalemail struct {
}

func (s Scalemail) Slot() string {
    return ""
}

func (s Scalemail) Use() {
}

func (s Scalemail) Save() string {
    return "Scalemail"
}

func (s Scalemail) PrettyPrint() string {
    return "Scalemail"
}
