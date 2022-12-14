package items

type Clothes struct {
    Quality string
}

func (c *Clothes) Slot() string {
    return "Clothes"
}

func (c *Clothes) Use() (string, []int) {
    return c.Action(), []int{}
}

func (c *Clothes) Save() string {
    return "Clothes," + c.Quality
}

func (c *Clothes) PrettyPrint() string {
    if c.Quality == "Costume" {
        return c.Quality
    } else {
        return c.Quality + " Clothes"
    }
}

func (c *Clothes) Function() string {
    return "clothes"
}

func (c *Clothes) Damage() (int, int, string) {
    return 0, 0, ""
}

func (c *Clothes) Action() string {
    return ""
}

func (c *Clothes) GetQuantity() int {
    return 1
}

func (c *Clothes) GetRange() []float64 {
    return []float64{0, 0}
}
