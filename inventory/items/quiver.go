package items

import (
    "fmt"
    "strconv"
)

type Quiver struct {
    Arrows int
}

func (q *Quiver) Slot() string {
    return "Torso"
}

func (q *Quiver) Use() (string, []int) {
    return q.Action(), []int{}
}

func (q *Quiver) Save() string {
    return "Quiver," + strconv.Itoa(q.Arrows)
}

func (q *Quiver) PrettyPrint() string {
    return fmt.Sprintf("Quiver (%d)", q.Arrows)
}

func (q *Quiver) Function() string {
    return "ammo"
}

func (q *Quiver) Damage() (int, int, string) {
    return 0, 0, ""
}

func (q *Quiver) Action() string {
    return ""
}

func (q *Quiver) GetQuantity() int {
    return q.Arrows
}

func (q *Quiver) GetRange() []float64 {
    return []float64{0, 0}
}
