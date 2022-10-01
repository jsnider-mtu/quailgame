package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TensersFloatingDisk struct {}

func (t TensersFloatingDisk) PrettyPrint() string {
    return "Tenser's Floating Disk"
}

func (t TensersFloatingDisk) GetLevel() int {
    return 1
}

func (a TensersFloatingDisk) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Tenser's Floating Disk is not implemented yet")
}
