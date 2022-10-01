package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ArmorOfAgathys struct {}

func (a ArmorOfAgathys) Cast(target *npcs.NPC) bool {
    log.Println("The spell Armor of Agathys is not implemented yet")
}

func (a ArmorOfAgathys) PrettyPrint() string {
    return "Armor of Agathys"
}

func (a ArmorOfAgathys) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Armor of Agathys is not implemented yet")
}
