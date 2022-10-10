package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ProtectionFromGoodAndEvil struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (p ProtectionFromGoodAndEvil) PrettyPrint() string {
    return "Protection from Good and Evil"
}

func (p ProtectionFromGoodAndEvil) GetLevel() int {
    return p.level
}

func (p ProtectionFromGoodAndEvil) GetCastTime() string {
    return p.casttime
}

func (p ProtectionFromGoodAndEvil) GetRange() int {
    return p.spellrange
}

func (p ProtectionFromGoodAndEvil) GetComponents() string {
    return p.components
}

func (p ProtectionFromGoodAndEvil) GetDuration() int {
    return p.duration
}

func (p ProtectionFromGoodAndEvil) GetConcentration() bool {
    return p.concentration
}

func (p ProtectionFromGoodAndEvil) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Protection from Good and Evil is not implemented yet")
}
