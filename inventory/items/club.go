package items

type Club struct {
}

func (c *Club) Slot() string {
    return "RightHand"
}

func (c *Club) Use() (string, []int) {
    return c.Action(), []int{}
}

func (c *Club) Save() string {
    return "Club"
}

func (c *Club) PrettyPrint() string {
    return "Club"
}

func (c *Club) Function() string {
    return "melee-light"
}

func (c *Club) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (c *Club) Action() string {
    return ""
}

func (c *Club) GetQuantity() int {
    return 1
}

func (c *Club) GetRange() []float64 {
    return []float64{0, 0}
}
