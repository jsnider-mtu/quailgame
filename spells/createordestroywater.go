package spells

type CreateOrDestroyWater struct {}

func (c CreateOrDestroyWater) Cast(target string) bool {
    log.Println("The spell Create Or Destroy Water is not implemented yet")
}

func (c CreateOrDestroyWater) PrettyPrint() string {
    return "Create or Destroy Water"
}
