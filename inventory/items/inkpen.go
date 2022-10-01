package items

type InkPen struct {
}

func (i InkPen) Slot() string {
    return "RightHand"
}

func (i InkPen) Use() {
}

func (i InkPen) Save() string {
    return "InkPen"
}

func (i InkPen) PrettyPrint() string {
    return "Ink Pen"
}
