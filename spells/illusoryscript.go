package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type IllusoryScript struct {}

func (i IllusoryScript) Cast(target *npcs.NPC) bool {
    log.Println("The spell Illusory Script is not implemented yet")
}

func (i IllusoryScript) PrettyPrint() string {
    return "Illusory Script"
}

func (a IllusoryScript) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Illusory Script is not implemented yet")
}
