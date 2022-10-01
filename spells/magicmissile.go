package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type MagicMissile struct {}

func (m MagicMissile) Cast(target *npcs.NPC) bool {
    log.Println("The spell Magic Missile is not implemented yet")
}

func (m MagicMissile) PrettyPrint() string {
    return "Magic Missile"
}
