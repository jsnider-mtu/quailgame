package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type IllusoryScript struct {}

func (i IllusoryScript) Cast(target *npcs.NPC) bool {
    log.Println("The spell Illusory Script is not implemented yet")
}

func (i IllusoryScript) PrettyPrint() string {
    return "Illusory Script"
}
