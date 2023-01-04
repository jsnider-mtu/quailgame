package items

type Chainmail struct {
}

func (c *Chainmail) Slot() string {
    return "Armor"
}

func (c *Chainmail) Use() (string, []int) {
    return "", []int{}
}

func (c *Chainmail) Save() string {
    return "Chainmail"
}

func (c *Chainmail) PrettyPrint() string {
    return "Chainmail"
}

func (c *Chainmail) Function() string {
    return "armor"
}

func (c *Chainmail) Damage() (int, int, string) {
    return 0, 0, ""
}

func (c *Chainmail) Action() string {
    return ""
}

func (c *Chainmail) GetQuantity() int {
    return 1
}

func (c *Chainmail) GetRange() []float64 {
    return []float64{0, 0}
}
