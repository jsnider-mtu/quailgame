package spells

type ExpeditiousRetreat struct {}

func (e ExpeditiousRetreat) Cast(target string) bool {
    log.Println("The spell Expeditious Retreat is not implemented yet")
}

func (e ExpeditiousRetreat) PrettyPrint() string {
    return "Expeditious Retreat"
}
