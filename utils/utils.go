package utils

import (
    "log"

    "golang.org/x/image/font"
    "golang.org/x/image/font/gofont/gomonobold"

    "github.com/golang/freetype/truetype"

)

func Fo() font.Face {
    fon, err := truetype.Parse(gomonobold.TTF)
    if err != nil {
        log.Fatal(err)
    }
    return truetype.NewFace(fon, &truetype.Options{Size: 20})
}
