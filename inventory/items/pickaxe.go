package items

import (
    "fmt"
)

type Pickaxe struct {
    Material string
    Durability int
}

func (p Pickaxe) Use() {
}

func (p Pickaxe) Save() string {
    return "Pickaxe," + p.Material + "," + fmt.Sprint(p.Durability)
}
