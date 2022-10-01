package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type TensersFloatingDisk struct {}

func (t TensersFloatingDisk) Cast(target *npcs.NPC) bool {
    log.Println("The spell Tenser's Floating Disk is not implemented yet")
}

func (t TensersFloatingDisk) PrettyPrint() string {
    return "Tenser's Floating Disk"
}
