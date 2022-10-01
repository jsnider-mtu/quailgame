package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Alarm struct {}

func (a Alarm) PrettyPrint() string {
    return "Alarm"
}

func (a Alarm) GetLevel() int {
    return 0
}

func (a Alarm) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Alarm is not implemented yet")
}
