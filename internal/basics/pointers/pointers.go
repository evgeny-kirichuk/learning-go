package pointers

import "fmt"

type Player struct {
	Name string
	HP int
}

func TakeDamageWithoutPointer(p Player, dmg int) {
	p.HP -= dmg;
	fmt.Printf("%s took %d damage\n", p.Name, dmg)
}

func TakeDamageWithPointer(p *Player, dmg int) {
	p.HP -= dmg;
	fmt.Printf("%s took %d damage\n", p.Name, dmg)
}

// function receiver
func (p *Player) TakeDamage(dmg int) {
	if p == nil {
		panic("player is nil")
	}
	p.HP -= dmg;
	fmt.Printf("%s took %d damage\n", p.Name, dmg)
	fmt.Printf("%s has %d/100HP remaining\n", p.Name, p.HP)
}

func RunPointers() {
	player := &Player{
		Name: "Evgeny",
		HP: 100,
	}

	// without pointer
	TakeDamageWithoutPointer(*player, 10)
	fmt.Printf("%s has %d/100HP remaining\n", player.Name, player.HP)

	// with pointer
	TakeDamageWithPointer(player, 10)
	fmt.Printf("%s has %d/100HP remaining\n", player.Name, player.HP)

	// with function receiver
	player.TakeDamage(10)
	player.TakeDamage(10)
}