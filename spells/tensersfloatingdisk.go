package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type TensersFloatingDisk struct {}

func (t TensersFloatingDisk) Cast(target *npcs.NPC) bool {
    log.Println("The spell Tenser's Floating Disk is not implemented yet")
}

func (t TensersFloatingDisk) PrettyPrint() string {
    return "Tenser's Floating Disk"
}

func (a TensersFloatingDisk) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Tenser's Floating Disk is not implemented yet")
}
