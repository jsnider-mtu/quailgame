package items

import (
    "fmt"
    "strconv"
)

type Quiver struct {
    Arrows int
}

func (q Quiver) Slot() string {
    return "Torso"
}

func (q Quiver) Use() {
}

func (q Quiver) Save() string {
    return "Quiver," + strconv.Itoa(q.Arrows)
}

func (q Quiver) PrettyPrint() string {
    return fmt.Sprintf("Quiver%16s", "Arrows: " + strconv.Itoa(q.Arrows))
}
