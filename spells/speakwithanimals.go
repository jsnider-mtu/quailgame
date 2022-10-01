package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type SpeakWithAnimals struct {}

func (s SpeakWithAnimals) Cast(target *npcs.NPC) bool {
    log.Println("The spell Speak with Animals is not implemented yet")
}

func (s SpeakWithAnimals) PrettyPrint() string {
    return "Speak with Animals"
}

func (a SpeakWithAnimals) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Speak with Animals is not implemented yet")
}
