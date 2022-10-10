package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type EldritchBlast struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (e EldritchBlast) PrettyPrint() string {
    return "Eldritch Blast"
}

func (e EldritchBlast) GetLevel() int {
    return e.level
}

func (e EldritchBlast) GetCastTime() string {
    return e.casttime
}

func (e EldritchBlast) GetRange() int {
    return e.spellrange
}

func (e EldritchBlast) GetComponents() string {
    return e.components
}

func (e EldritchBlast) GetDuration() int {
    return e.duration
}

func (e EldritchBlast) GetConcentration() bool {
    return e.concentration
}

func (e EldritchBlast) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Eldritch Blast is not implemented yet")
}
