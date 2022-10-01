package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type DetectEvilAndGood struct {}

func (d DetectEvilAndGood) Cast(target *npcs.NPC) bool {
    log.Println("The spell Detect Evil and Good is not implemented yet")
}

func (d DetectEvilAndGood) PrettyPrint() string {
    return "Detect Evil and Good"
}

func (a DetectEvilAndGood) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Detect Evil and Good is not implemented yet")
}
