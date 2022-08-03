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
    cscount int = 0
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
            cscount++
        }
        if cscount > 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
            cscount = len(textstr)
        }
        if cscount < len(textstr) {
            text.Draw(screen, textstr[:cscount], *fo, 64, 64, color.White)
            return false
        } else if cscount > len(textstr) + 20 {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            cscount = 0
            return true
        } else {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            return false
        }
    case 1:
        textstr := "The Quail Kingdom...\n\n"+
                   "Second cutscene."
        if count % 5 == 0 {
            cscount++
        }
        if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
            cscount = len(textstr)
        }
        if cscount < len(textstr) {
            text.Draw(screen, textstr[:cscount], *fo, 64, 64, color.White)
            return false
        } else if cscount > len(textstr) + 20 {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            cscount = 0
            return true
        } else {
            text.Draw(screen, textstr, *fo, 64, 64, color.White)
            return false
        }
    default:
        return true
    }
}
