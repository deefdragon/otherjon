package main

import "fmt"

type playerData struct {
	name        string
	level       int
	xp          int
	Health      int
	Mana        int
	Stamina     int
	maxHealth   int
	maxMana     int
	maxStamina  int
	xpTillLevel int
	gold        int
}

// Method for health potion, increases current health up to a maximum of max health
func (p *playerData) kevJuice(Health int, Mana int, Stamina int) {
	p.Health += 100
	p.Mana += 100
	p.Stamina += 50
	fmt.Printf("You drank some kevJuice! You feel revitalized")
	if p.Health > p.maxHealth {
		p.Health = p.maxHealth
	}
	if p.Stamina > p.maxStamina {
		p.Stamina = p.maxStamina
	}
	if p.Mana > p.maxMana {
		p.Mana = p.maxMana
	}
}

// Method for decreasing the player's health if an enemy's swing connects
func (p *playerData) gotHurt(Health int) {
	hurt := true
	if hurt {
		p.Health -= 25
		fmt.Printf("Ouch! You have %d health remaining\n", p.Health)
	} else {
		fmt.Printf("A near miss, you're unscathed!\n")
	}
}

// Method for casting fireball, if successful, will cost mana
func (p *playerData) fireball(Mana int) {
	fizzle := false
	if fizzle {
		fmt.Printf("Oh no! Your spell fizzled!\n")
	} else {
		p.Mana -= 50
		fmt.Printf("BOOM! Your enemy is burning, you have %d mana remaining\n", p.Mana)
	}
}

// Method for a Regular swing, costing stamina hit or miss
func (p *playerData) swing(Stamina int) {
	hit := true
	p.Stamina -= 10
	if hit {
		fmt.Printf("A mighty strike! You have %d stamina left\n", p.Stamina)
	} else {
		fmt.Printf("You whiffed! You have %d stamina left\n", p.Stamina)
	}
}

// Method that levels the player up, increasing their level count and stats
func (p *playerData) levelUp() {
	p.xp = 0
	p.xpTillLevel += 50
	fmt.Printf("You leveled up! You gained 50 maxHP, 50 maxMP, and 10 maxSP\n")
	p.maxHealth += 50
	p.maxMana += 50
	p.maxStamina += 10
	p.level += 1
}

// Method to list off player1's stats
func (p *playerData) pStatus() {
	fmt.Printf("Character: %s\n", p.name)
	fmt.Printf("Level: %d\n", p.level)
	fmt.Printf("Experience: %d out of %d\n", p.xp, p.xpTillLevel)
	fmt.Printf("HP: %d out of %d\n", p.Health, p.maxHealth)
	fmt.Printf("MP: %d out of %d\n", p.Mana, p.maxMana)
	fmt.Printf("SP: %d out of %d\n", p.Stamina, p.maxStamina)
}

// Method that rewards the player when a monster is defeated
func (p *playerData) killedMonster(xp int, gold int) {
	victory := true
	if victory {
		p.xp += 100
		p.gold += 100
		if p.xp >= p.xpTillLevel {
			p.levelUp()
		}
	}
}
func main() {
	var player1 = playerData{"Jon", 1, 0, 100, 100, 50, 100, 100, 50, 100, 0}
	player1.pStatus()
	player1.gotHurt(player1.Health)
	player1.fireball(player1.Mana)
	player1.fireball(player1.Mana)
	player1.swing(player1.Stamina)
	player1.killedMonster(player1.xp, player1.gold)
	player1.kevJuice(player1.Health, player1.Mana, player1.Stamina)
	player1.pStatus()
}
