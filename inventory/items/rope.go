package items

import (
    "fmt"
    "strconv"
)

type Rope struct {
    Length int
}

func (r Rope) Slot() string {
    return "BothHands"
}

func (r Rope) Use() {
}

func (r Rope) Save() string {
    return "Rope," + strconv.Itoa(r.Length)
}

func (r Rope) PrettyPrint() string {
    return fmt.Sprintf("Rope (%d)", r.Length)
}
