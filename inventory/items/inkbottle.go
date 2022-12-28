package items

type InkBottle struct {
}

func (i InkBottle) Slot() string {
    return ""
}

func (i InkBottle) Use() (string, []int) {
}

func (i InkBottle) Save() string {
    return "InkBottle"
}

func (i InkBottle) PrettyPrint() string {
    return "InkBottle"
}

func (i InkBottle) Function() string {
    return "writing"
}

func (i InkBottle) Damage() (int, int, string) {
    return 0, 0, ""
}
