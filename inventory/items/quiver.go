package items

import "strconv"

type Quiver struct {
    Arrows int
}

func (q Quiver) Slot() string {
    return ""
}

func (q Quiver) Use() {
}

func (q Quiver) Save() string {
    return "Quiver," + strconv.Itoa(q.Arrows)
}
