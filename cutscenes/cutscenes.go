package cutscenes

import (
//    "fmt"
    "image/color"

    "golang.org/x/image/font"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"
)

var (
    f int = 0
)

func CutScene(screen *ebiten.Image, cs, count int, fo *font.Face) bool {
    switch cs {
    case 0:
        textstr := "The Quail Kingdom...\n\n"+
                   "A safe respite for many. A place of deep history,\n"+
                   "a place of ancient warfare, but now a place of calm\n"+
                   "peace. Your travels have brought you to the capital\n"+
                   "city on the eve of the Festival of the Watermelon.\n"+
                   "While exploring you find an old man searching for\n"+
                   "a book. Although you wish him well your priority now\n"+
                   "is ale and hedonism."
        if count % 5 == 0 {
            f++
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            f = len(textstr)
        }
        if f < len(textstr) {
            text.Draw(screen, textstr[:f], *fo, 64, 64, color.White)
            return false
        } else if f > len(textstr) + 20 {
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
