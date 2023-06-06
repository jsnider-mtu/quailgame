package main

import (
    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/utils"
)

func main() {
    utils.Byteslice2file(assets.KingQuail_PNG, "kingquail.png")
}
