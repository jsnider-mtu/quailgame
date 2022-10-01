package spells

type TashasHideousLaughter struct {}

func (t TashasHideousLaughter) Cast(target string) bool {
    log.Println("The spell Tasha's Hideous Laughter is not implemented yet")
}

func (t TashasHideousLaughter) PrettyPrint() string {
    return "Tasha's Hideous Laughter"
}
