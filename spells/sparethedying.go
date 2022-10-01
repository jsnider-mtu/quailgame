package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type SpareTheDying struct {}

func (s SpareTheDying) Cast(target *npcs.NPC) bool {
    log.Println("The spell Spare the Dying is not implemented yet")
}

func (s SpareTheDying) PrettyPrint() string {
    return "Spare the Dying"
}
