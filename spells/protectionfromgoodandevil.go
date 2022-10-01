package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ProtectionFromGoodAndEvil struct {}

func (p ProtectionFromGoodAndEvil) Cast(target *npcs.NPC) bool {
    log.Println("The spell Protection from Good and Evil is not implemented yet")
}

func (p ProtectionFromGoodAndEvil) PrettyPrint() string {
    return "Protection from Good and Evil"
}
