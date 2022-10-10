package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Message struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (m Message) PrettyPrint() string {
    return "Message"
}

func (m Message) GetLevel() int {
    return m.level
}

func (m Message) GetCastTime() string {
    return m.casttime
}

func (m Message) GetRange() int {
    return m.spellrange
}

func (m Message) GetComponents() string {
    return m.components
}

func (m Message) GetDuration() int {
    return m.duration
}

func (m Message) GetConcentration() bool {
    return m.concentration
}

func (m Message) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Message is not implemented yet")
}
