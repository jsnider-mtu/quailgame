package items

import (
    "fmt"
    "strconv"
)

type Perfume struct {
    Quantity int
}

func (p *Perfume) Slot() string {
    return ""
}

func (p *Perfume) Use() (string, []int) {
    return p.Action(), []int{2, 600}
}

func (p *Perfume) Save() string {
    return "Perfume," + strconv.Itoa(p.Quantity)
}

func (p *Perfume) PrettyPrint() string {
    return fmt.Sprintf("Perfume (%d)", p.Quantity)
}

func (p *Perfume) Function() string {
    return "charisma-buff"
}

func (p *Perfume) Damage() (int, int, string) {
    return 0, 0, ""
}

func (p *Perfume) Action() string {
    return "charisma"
}

func (p *Perfume) GetQuantity() int {
    return p.Quantity
}

func (p *Perfume) GetRange() []float64 {
    return []float64{0, 0}
}
