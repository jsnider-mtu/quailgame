package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FireBolt struct {}

func (f FireBolt) Cast(target *npcs.NPC) bool {
    log.Println("The spell Fire Bolt is not implemented yet")
}

func (f FireBolt) PrettyPrint() string {
    return "Fire Bolt"
}
