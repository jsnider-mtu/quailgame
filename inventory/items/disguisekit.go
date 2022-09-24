package items

type Disguisekit struct {
}

func (d Disguisekit) Slot() string {
    return ""
}

func (d Disguisekit) Use() {
}

func (d Disguisekit) Save() string {
    return "Disguisekit"
}

func (d Disguisekit) PrettyPrint() string {
    return "Disguisekit"
}
