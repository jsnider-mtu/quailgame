package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type AcidSplash struct {}

func (a AcidSplash) Cast(target *npcs.NPC) bool {
    log.Println("The spell Acid Splash is not implemented yet")
}

func (a AcidSplash) PrettyPrint() string {
    return "Acid Splash"
}
