package spells

type ColorSpray struct {}

func (c ColorSpray) Cast(target string) bool {
    log.Println("The spell Color Spray is not implemented yet")
}

func (c ColorSpray) PrettyPrint() string {
    return "Color Spray"
}
