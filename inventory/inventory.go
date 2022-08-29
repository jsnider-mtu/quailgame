package inventory

type Item interface {
    Use()
    Save() string
}

type Inv struct {
    Cap int
    Items []Item
}

func (i *Inv) Drop(item Item) {
    for a, b := range i.Items {
        if b == item {
            i.Items = append(i.Items[:a], i.Items[a + 1:]...)
            break
        }
    }
}
