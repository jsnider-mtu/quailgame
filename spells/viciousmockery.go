package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ViciousMockery struct {}

func (v ViciousMockery) PrettyPrint() string {
    return "Vicious Mockery"
}

func (v ViciousMockery) GetLevel() int {
    return 0
}

func (a ViciousMockery) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Vicious Mockery is not implemented yet")
}
