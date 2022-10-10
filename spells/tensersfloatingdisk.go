package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TensersFloatingDisk struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (t TensersFloatingDisk) PrettyPrint() string {
    return "Tenser's Floating Disk"
}

func (t TensersFloatingDisk) GetLevel() int {
    return t.level
}

func (t TensersFloatingDisk) GetCastTime() string {
    return t.casttime
}

func (t TensersFloatingDisk) GetRange() int {
    return t.range
}

func (t TensersFloatingDisk) GetComponents() string {
    return t.components
}

func (t TensersFloatingDisk) GetDuration() int {
    return t.duration
}

func (t TensersFloatingDisk) GetConcentration() bool {
    return t.concentration
}

func (t TensersFloatingDisk) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Tenser's Floating Disk is not implemented yet")
}
