package spells

type DivineFavor struct {}

func (d DivineFavor) Cast(target string) bool {
    log.Println("The spell Divine Favor is not implemented yet")
}

func (d DivineFavor) PrettyPrint() string {
    return "Divine Favor"
}
