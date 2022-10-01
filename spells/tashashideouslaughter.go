package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type TashasHideousLaughter struct {}

func (t TashasHideousLaughter) Cast(target *npcs.NPC) bool {
    log.Println("The spell Tasha's Hideous Laughter is not implemented yet")
}

func (t TashasHideousLaughter) PrettyPrint() string {
    return "Tasha's Hideous Laughter"
}
