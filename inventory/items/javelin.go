package items

import (
    "fmt"
    "strconv"
)

type Javelin struct {
    Quantity int
}

func (j *Javelin) Slot() string {
    return "RightHand"
}

func (j *Javelin) Use() (string, []int) {
    return j.Action(), []int{144, 576}
}

func (j *Javelin) Save() string {
    return "Javelin," + strconv.Itoa(j.Quantity)
}

func (j *Javelin) PrettyPrint() string {
    return fmt.Sprintf("Javelin (%d)", j.Quantity)
}

func (j *Javelin) Function() string {
    return "melee-throw"
}

func (j *Javelin) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (j *Javelin) Action() string {
    return "throw"
}

func (j *Javelin) GetQuantity() int {
    return 1
}

func (j *Javelin) GetRange() []float64 {
    return []float64{144.0, 576.0}
}
