package spells

type ArmorOfAgathys struct {}

func (a ArmorOfAgathys) Cast(target string) bool {
    log.Println("The spell Armor of Agathys is not implemented yet")
}

func (a ArmorOfAgathys) PrettyPrint() string {
    return "Armor of Agathys"
}
