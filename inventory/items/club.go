package items

type Club struct {
}

func (c *Club) Slot() string {
    return "RightHand"
}

func (c *Club) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (c *Club) Save() string {
    return "Club"
}

func (c *Club) PrettyPrint() string {
    return "Club"
}

func (c *Club) Function() string {
    return "melee"
}

func (c *Club) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}

func (c *Club) Action() string {
    return ""
}
