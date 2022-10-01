package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type AnimalFriendship struct {}

func (a AnimalFriendship) Cast(target *npcs.NPC) bool {
    log.Println("The spell Animal Friendship is not implemented yet")
}

func (a AnimalFriendship) PrettyPrint() string {
    return "Animal Friendship"
}
