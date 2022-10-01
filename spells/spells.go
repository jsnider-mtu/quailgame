package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Spell interface {
    Cast(*npcs.NPC) bool
    PrettyPrint() string
}
