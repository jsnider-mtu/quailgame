package spells

type FogCloud struct {}

func (f FogCloud) Cast(target string) bool {
    log.Println("The spell Fog Cloud is not implemented yet")
}

func (f FogCloud) PrettyPrint() string {
    return "Fog Cloud"
}
