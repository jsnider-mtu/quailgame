package cutscenes

import (
    "bytes"
//    "fmt"
    "log"
    "image"
    "image/color"
    _ "image/png"

    "golang.org/x/image/font"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/shaders"
)

var (
    cscount int = 0
    pic0Image *ebiten.Image
    pic2Image *ebiten.Image
)

var textstrs = make([]string, 0)
var picsarr = make([]*ebiten.Image, 0)

func CutScene(screen *ebiten.Image, cs, count int, fo *font.Face) bool {
    if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        cscount = 0
        if len(textstrs) > 1 {
            textstrs = textstrs[1:]
            if len(picsarr) > 1 {
                picsarr = picsarr[1:]
            }
            return false
        } else {
            textstrs = make([]string, 0)
            picsarr = make([]*ebiten.Image, 0)
            return true
        }
    }
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
        if len(picsarr) == 0 {
            pic0image, _, err := image.Decode(bytes.NewReader(assets.OldMan_PNG))
            if err != nil {
                log.Fatal(err)
            }
            pic0Image = ebiten.NewImageFromImage(pic0image)
            pic1Image := ebiten.NewImage(300, 300)
            pic2image, _, err := image.Decode(bytes.NewReader(assets.KingQuail_PNG))
            if err != nil {
                log.Fatal(err)
            }
            pic2Image = ebiten.NewImageFromImage(pic2image)
            picsarr = append(picsarr, pic0Image, pic1Image, pic2Image)
        }
        if count % 3 == 0 {
            cscount++
        }
        //if cscount > 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        //    cscount = len(textstr)
        //}
        cx, cy := ebiten.CursorPosition()
        whiteImage := ebiten.NewImage(300, 300)
        //blackImage := ebiten.NewImage(300, 300)
        //blackImage.Fill(color.Black)
        whiteImage.Fill(color.White)
        s, err := ebiten.NewShader([]byte(shaders.Lighting_go))
        if err != nil {
            log.Fatal(err)
        }
        sgm := ebiten.GeoM{}
        sgm.Translate(float64(234), float64(276))
        sop := &ebiten.DrawRectShaderOptions{GeoM: sgm}
        if count > 220 {
            sop.Uniforms = map[string]interface{}{
                "Time": float32(count - 220) / 60,
                "Cursor": []float32{float32(cx), float32(cy)},
                "ScreenSize": []float32{float32(768), float32(576)},
            }
        } else {
            sop.Uniforms = map[string]interface{}{
                "Time": float32(0),
                "Cursor": []float32{float32(cx), float32(cy)},
                "ScreenSize": []float32{float32(768), float32(576)},
            }
        }
        sop.Images[0] = picsarr[0]
        sop.Images[1] = whiteImage
        sop.Images[2] = whiteImage
        sop.Images[3] = whiteImage
        i := float64(cscount) / float64(len(textstrs[0]) + 50)
        if cscount < len(textstrs[0]) {
            text.Draw(screen, textstrs[0][:cscount], *fo, 64, 64, color.White)
            pgm := ebiten.GeoM{}
            pgm.Translate(float64(234), float64(276))
            pcm := ebiten.ColorM{}
            pcm.Scale(1.0, 1.0, 1.0, i)
            screen.DrawImage(picsarr[0], &ebiten.DrawImageOptions{GeoM: pgm, ColorM: pcm})
            screen.DrawRectShader(300, 300, s, sop)
            return false
        } else if cscount > len(textstrs[0]) + 50 {
            text.Draw(screen, textstrs[0], *fo, 64, 64, color.White)
            pgm := ebiten.GeoM{}
            pgm.Translate(float64(234), float64(276))
            pcm := ebiten.ColorM{}
            pcm.Scale(1.0, 1.0, 1.0, i)
            screen.DrawImage(picsarr[0], &ebiten.DrawImageOptions{GeoM: pgm, ColorM: pcm})
            screen.DrawRectShader(300, 300, s, sop)
            cscount = 0
            if len(picsarr) > 1 {
                picsarr = picsarr[1:]
            }
            if len(textstrs) > 1 {
                textstrs = textstrs[1:]
                return false
            } else {
                textstrs = make([]string, 0)
                picsarr = make([]*ebiten.Image, 0)
                return true
            }
        } else {
            text.Draw(screen, textstrs[0], *fo, 64, 64, color.White)
            pgm := ebiten.GeoM{}
            pgm.Translate(float64(234), float64(276))
            pcm := ebiten.ColorM{}
            pcm.Scale(1.0, 1.0, 1.0, i)
            screen.DrawImage(picsarr[0], &ebiten.DrawImageOptions{GeoM: pgm, ColorM: pcm})
            screen.DrawRectShader(300, 300, s, sop)
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
