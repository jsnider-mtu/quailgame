package items

type Scalemail struct {
}

func (s Scalemail) Slot() string {
    return "Armor"
}

func (s Scalemail) Use() {
}

func (s Scalemail) Save() string {
    return "Scalemail"
}

func (s Scalemail) PrettyPrint() string {
    return "Scalemail"
}

func (s Scalemail) Function() string {
    return "armor"
}

func (s Scalemail) Damage() (int, int, string) {
    return 0, 0, ""
}
