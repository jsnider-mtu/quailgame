package items

type Net struct {
}

func (n *Net) Slot() string {
    return "RightHand"
}

func (n *Net) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (n *Net) Save() string {
    return "Net"
}

func (n *Net) PrettyPrint() string {
    return "Net"
}

func (n *Net) Function() string {
    return "range"
}

func (n *Net) Damage() (int, int, string) {
    return 0, 0, ""
}

func (n *Net) Action() string {
    return ""
}

func (n *Net) GetQuantity() int {
    return 1
}
