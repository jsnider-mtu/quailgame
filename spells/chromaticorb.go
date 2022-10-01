package spells

type ChromaticOrb struct {}

func (c ChromaticOrb) Cast(target string) bool {
    log.Println("The spell Chromatic Orb is not implemented yet")
}

func (c ChromaticOrb) PrettyPrint() string {
    return "Chromatic Orb"
}
