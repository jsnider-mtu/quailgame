package items

type InkPen struct {
}

func (i InkPen) Slot() string {
    return "RightHand"
}

func (i InkPen) Use() (string, []int) {
    return "write", []int{}
}

func (i InkPen) Save() string {
    return "InkPen"
}

func (i InkPen) PrettyPrint() string {
    return "Ink Pen"
}

func (i InkPen) Function() string {
    return "writing"
}

func (i InkPen) Damage() (int, int, string) {
    return 0, 0, ""
}
