package spells

type AcidSplash struct {}

func (a AcidSplash) Cast(target string) bool {
    log.Println("The spell Acid Splash is not implemented yet")
}

func (a AcidSplash) PrettyPrint() string {
    return "Acid Splash"
}
