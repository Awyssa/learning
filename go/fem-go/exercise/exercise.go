package exercise

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	// Add code here
}

func (p *Player) DropItem(itemName string) {
	// Add code here
}

func (p *Player) UseItem(itemName string) {
	// add code here
}
