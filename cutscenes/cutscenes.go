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

var textstrs = make([]string, 0)

func CutScene(screen *ebiten.Image, cs, count int, fo *font.Face) bool {
    switch cs {
    case 0:
        if len(textstrs) == 0 {
            textstr := "The Quail Kingdom...\n\n"+
                       "A safe respite for many. A place of deep history,\n"+
                       "a place of ancient warfare, but now a place of calm\n"+
                       "peace. Your travels have brought you to the capital\n"+
                       "city on the eve of the Festival of the Watermelon.\n"+
                       "While exploring you find an old man searching for\n"+
                       "a book. Although you wish him well your priority now\n"+
                       "is ale and hedonism."
            textstr2 := "But later that evening, by chance you find this book...\n"+
                        "It’s very old and written in a language you cannot\n"+
                        "decipher. The next day you seek to find the old man and\n"+
                        "return his trinket. Instead, you find him murdered!\n"+
                        "Fearing for your life, you cast the book into a river\n"+
                        "and head back to your room, only to find the book laying\n"+
                        "on your bed soaked with the stream's fresh water. The news\n"+
                        "of the murder has now reached the Lords."
            textstr3 := "Discover the meaning of this magical relic and its\n"+
                        "prophecies. The quest will lay bare the tale of\n"+
                        "brother against brother, the sundering of a kingdom,\n"+
                        "and the rising from the ashes\n\n"+
                        "-- a new sovereign -- King Quail."
            textstrs = append(textstrs, textstr, textstr2, textstr3)
        }
        if count % 3 == 0 {
            cscount++
        }
        //if cscount > 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        //    cscount = len(textstr)
        //}
        if cscount < len(textstrs[0]) {
            text.Draw(screen, textstrs[0][:cscount], *fo, 64, 64, color.White)
            return false
        } else if cscount > len(textstrs[0]) + 50 {
            text.Draw(screen, textstrs[0], *fo, 64, 64, color.White)
            cscount = 0
            if len(textstrs) > 1 {
                textstrs = textstrs[1:]
                return false
            } else {
                textstrs = make([]string, 0)
                return true
            }
        } else {
            text.Draw(screen, textstrs[0], *fo, 64, 64, color.White)
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
