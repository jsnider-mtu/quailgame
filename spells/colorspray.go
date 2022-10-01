package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ColorSpray struct {}

func (c ColorSpray) Cast(target *npcs.NPC) bool {
    log.Println("The spell Color Spray is not implemented yet")
}

func (c ColorSpray) PrettyPrint() string {
    return "Color Spray"
}
