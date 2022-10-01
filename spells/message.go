package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Message struct {}

func (m Message) Cast(target *npcs.NPC) bool {
    log.Println("The spell Message is not implemented yet")
}

func (m Message) PrettyPrint() string {
    return "Message"
}

func (a Message) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Message is not implemented yet")
}
