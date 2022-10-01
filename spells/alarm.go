package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Alarm struct {}

func (a Alarm) Cast(target *npcs.NPC) bool {
    log.Println("The spell Alarm is not implemented yet")
}

func (a Alarm) PrettyPrint() string {
    return "Alarm"
}

func (a Alarm) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Alarm is not implemented yet")
}
