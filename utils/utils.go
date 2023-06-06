package utils

import (
    "log"
    "os"

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

func Byteslice2file(bs []byte, path string) {
    f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    _, err = f.Write(bs)
    if err != nil {
        log.Fatal(err)
    }
    f.Close()
}
