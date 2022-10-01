package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type CharmPerson struct {}

func (c CharmPerson) Cast(target *npcs.NPC) bool {
    log.Println("The spell Charm Person is not implemented yet")
}

func (c CharmPerson) PrettyPrint() string {
    return "Charm Person"
}
