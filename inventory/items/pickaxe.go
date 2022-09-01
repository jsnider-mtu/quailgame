package items

import (
    "fmt"
)

type Pickaxe struct {
    material string
    durability int
}

func (p Pickaxe) GetMaterial() string {
    return p.material
}

func (p Pickaxe) GetDurability() int {
    return p.durability
}

func (p Pickaxe) Slot() string {
    return "RightHand"
}

func (p Pickaxe) Use() {
}

func (p Pickaxe) Save() string {
    return "Pickaxe," + p.material + "," + fmt.Sprint(p.durability)
}
