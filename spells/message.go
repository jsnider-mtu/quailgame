package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Message struct {}

func (m Message) PrettyPrint() string {
    return "Message"
}

func (m Message) GetLevel() int {
    return 0
}

func (a Message) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Message is not implemented yet")
}
