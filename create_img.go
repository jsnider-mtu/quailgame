// +build ignore

package main

import (
//    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/levels/lvlimages"
    "github.com/jsnider-mtu/quailgame/utils"
)

func main() {
    //utils.Byteslice2file(assets.KingQuail_PNG, "kingquail.png")
//    utils.Byteslice2file(assets.Overworld_PNG, "overworld.png")
    utils.Byteslice2file(lvlimages.Two_PNG, "lvltwo.png")
}
