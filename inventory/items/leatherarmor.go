package items

type LeatherArmor struct {
}

func (l LeatherArmor) Slot() string {
    return "Armor"
}

func (l LeatherArmor) Use() {
}

func (l LeatherArmor) Save() string {
    return "LeatherArmor"
}

func (l LeatherArmor) PrettyPrint() string {
    return "Leather Armor"
}
