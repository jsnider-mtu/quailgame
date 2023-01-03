package items

type InkPen struct {
}

func (i *InkPen) Slot() string {
    return ""
}

func (i *InkPen) Use() (string, []int) {
    return "write", []int{}
}

func (i *InkPen) Save() string {
    return "InkPen"
}

func (i *InkPen) PrettyPrint() string {
    return "Ink Pen"
}

func (i *InkPen) Function() string {
    return "writing"
}

func (i *InkPen) Damage() (int, int, string) {
    return 0, 0, ""
}

func (i *InkPen) Action() string {
    return "write"
}

func (i *InkPen) GetQuantity() int {
    return 1
}
