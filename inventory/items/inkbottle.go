package items

type InkBottle struct {
}

func (i *InkBottle) Slot() string {
    return ""
}

func (i *InkBottle) Use() (string, []int) {
    return "", []int{}
}

func (i *InkBottle) Save() string {
    return "InkBottle"
}

func (i *InkBottle) PrettyPrint() string {
    return "Ink Bottle"
}

func (i *InkBottle) Function() string {
    return "writing"
}

func (i *InkBottle) Damage() (int, int, string) {
    return 0, 0, ""
}

func (i *InkBottle) Action() string {
    return ""
}

func (i *InkBottle) GetQuantity() int {
    return 1
}

func (i *InkBottle) GetRange() []float64 {
    return []float64{0, 0}
}
