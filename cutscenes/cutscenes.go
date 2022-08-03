package cutscenes

import (
//    "fmt"
    "image/color"

    "golang.org/x/image/font"

    "github.com/hajimehoshi/ebiten/v2"
//    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"
)

var (
    f int = 0
)

func CutScene(screen *ebiten.Image, cs, count int, fo *font.Face) bool {
    switch cs {
    case 0:
        textstr := "test string to check out animations\nline no 2, more random text"
        if count % 8 == 0 {
            f++
        }
        if f < len(textstr) {
            text.Draw(screen, textstr[:f], *fo, 64, 64, color.White)
            return false
        } else if f > len(textstr) + 10 {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            f = 0
            return true
        } else {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            return false
        }
    default:
        return true
    }
}
