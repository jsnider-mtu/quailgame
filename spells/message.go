package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Message struct {}

func (m Message) Cast(target *npcs.NPC) bool {
    log.Println("The spell Message is not implemented yet")
}

func (m Message) PrettyPrint() string {
    return "Message"
}
