package items

type Chainmail struct {
}

func (c Chainmail) Slot() string {
    return ""
}

func (c Chainmail) Use() {
}

func (c Chainmail) Save() string {
    return "Chainmail"
}
