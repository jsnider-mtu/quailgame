package items

import (
    "fmt"
    "strconv"
)

type Handaxe struct {
    Quantity int
}

func (h *Handaxe) Slot() string {
    return "RightHand"
}

func (h *Handaxe) Use() (string, []int) {
    return h.Action(), []int{96, 288}
}

func (h *Handaxe) Save() string {
    return "Handaxe," + strconv.Itoa(h.Quantity)
}

func (h *Handaxe) PrettyPrint() string {
    return fmt.Sprintf("Handaxe (%d)", h.Quantity)
}

func (h *Handaxe) Function() string {
    return "melee-light-throw"
}

func (h *Handaxe) Damage() (int, int, string) {
    return 1, 6, "slashing"
}

func (h *Handaxe) Action() string {
    if h.Quantity > 0 {
        return "throw"
    }
    return ""
}

func (h *Handaxe) GetQuantity() int {
    return h.Quantity
}

func (h *Handaxe) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
