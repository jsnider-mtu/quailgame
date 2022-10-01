package items

type InkBottle struct {
}

func (i InkBottle) Slot() string {
    return ""
}

func (i InkBottle) Use() {
}

func (i InkBottle) Save() string {
    return "InkBottle"
}

func (i InkBottle) PrettyPrint() string {
    return "InkBottle"
}
