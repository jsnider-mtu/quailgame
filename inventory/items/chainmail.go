package items

type Chainmail struct {
}

func (c Chainmail) Slot() string {
    return "Armor"
}

func (c Chainmail) Use() {
}

func (c Chainmail) Save() string {
    return "Chainmail"
}

func (c Chainmail) PrettyPrint() string {
    return "Chainmail"
}

func (c Chainmail) Function() string {
    return "armor"
}

func (c Chainmail) Damage() (int, int, string) {
    return 0, 0, ""
}
