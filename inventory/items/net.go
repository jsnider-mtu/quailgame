package items

type Net struct {
}

func (n *Net) Slot() string {
    return "RightHand"
}

func (n *Net) Use() (string, []int) {
    return n.Action(), []int{24, 72}
}

func (n *Net) Save() string {
    return "Net"
}

func (n *Net) PrettyPrint() string {
    return "Net"
}

func (n *Net) Function() string {
    return "range-throw"
}

func (n *Net) Damage() (int, int, string) {
    return 0, 0, ""
}

func (n *Net) Action() string {
    return "throw"
}

func (n *Net) GetQuantity() int {
    return 1
}

func (n *Net) GetRange() []float64 {
    return []float64{24.0, 72.0}
}
