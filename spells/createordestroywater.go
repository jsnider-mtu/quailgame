package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CreateOrDestroyWater struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (c CreateOrDestroyWater) PrettyPrint() string {
    return "Create or Destroy Water"
}

func (c CreateOrDestroyWater) GetLevel() int {
    return c.level
}

func (c CreateOrDestroyWater) GetCastTime() string {
    return c.casttime
}

func (c CreateOrDestroyWater) GetRange() int {
    return c.spellrange
}

func (c CreateOrDestroyWater) GetComponents() string {
    return c.components
}

func (c CreateOrDestroyWater) GetDuration() int {
    return c.duration
}

func (c CreateOrDestroyWater) GetConcentration() bool {
    return c.concentration
}

func (c CreateOrDestroyWater) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Create or Destroy Water is not implemented yet")
}
