package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type DetectMagic struct {}

func (d DetectMagic) Cast(target *npcs.NPC) bool {
    log.Println("The spell Detect Magic is not implemented yet")
}

func (d DetectMagic) PrettyPrint() string {
    return "Detect Magic"
}
