package items

type DisguiseKit struct {
}

func (d DisguiseKit) Slot() string {
    return ""
}

func (d DisguiseKit) Use() {
}

func (d DisguiseKit) Save() string {
    return "DisguiseKit"
}

func (d DisguiseKit) PrettyPrint() string {
    return "Disguise Kit"
}

func (d DisguiseKit) Function() string {
    return "disguise"
}

func (d DisguiseKit) Damage() (int, int, string) {
    return 0, 0, ""
}
