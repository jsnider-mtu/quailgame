package items

import (
    "fmt"
    "strconv"
)

type Rope struct {
    Length int
}

func (r *Rope) Slot() string {
    return "BothHands"
}

func (r *Rope) Use() (string, []int) {
    return r.Action(), []int{}
}

func (r *Rope) Save() string {
    return "Rope," + strconv.Itoa(r.Length)
}

func (r *Rope) PrettyPrint() string {
    return fmt.Sprintf("Rope (%d)", r.Length)
}

func (r *Rope) Function() string {
    return "rope"
}

func (r *Rope) Damage() (int, int, string) {
    return 0, 0, ""
}

func (r *Rope) Action() string {
    return ""
}

func (r *Rope) GetQuantity() int {
    return r.Length
}

func (r *Rope) GetRange() []float64 {
    return []float64{0, 0}
}
