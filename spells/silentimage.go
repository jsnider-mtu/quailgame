package spells

type SilentImage struct {}

func (s SilentImage) Cast(target string) bool {
    log.Println("The spell Silent Image is not implemented yet")
}

func (s SilentImage) PrettyPrint() string {
    return "Silent Image"
}
