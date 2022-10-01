package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Alarm struct {}

func (a Alarm) Cast(target *npcs.NPC) bool {
    log.Println("The spell Alarm is not implemented yet")
}

func (a Alarm) PrettyPrint() string {
    return "Alarm"
}
