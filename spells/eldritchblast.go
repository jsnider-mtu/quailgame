package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type EldritchBlast struct {}

func (e EldritchBlast) Cast(target *npcs.NPC) bool {
    log.Println("The spell Eldritch Blast is not implemented yet")
}

func (e EldritchBlast) PrettyPrint() string {
    return "Eldritch Blast"
}

func (a EldritchBlast) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Eldritch Blast is not implemented yet")
}
