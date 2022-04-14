package player

import "fmt"

type Player struct {
	//TODO Getname,
	name string

	//TODO Getlevel,
	level int

	//TODO Getxp, AddXP
	xp int

	//TODO Gethealth, ChangeHealth, IsDead
	health int

	//TODO Getmana, ChangeMana, CanCast
	mana int

	//TODO Getstamina, ChangeStamina
	stamina int

	//TODO GetmaxHealth,
	maxHealth int

	//TODO GetmaxMana,
	maxMana int

	//TODO GetmaxStamina,
	maxStamina int

	//TODO GetxpTillLevel,
	xpTillLevel int

	//TODO Getgold, ChangeGold, HasGold
	gold int

	//TODO print player stats to console (Status)

	//TODO remove all unnecessary functions.
}

func New() *Player {

	return &Player{}
}

func (p *Player) GetGold() int {
	return p.gold
}

func (p *Player) changeHealth(health int) {

	if p.health <= health {
		p.health = 0
		return
	}
	p.health -= health
}

// Method for health potion, increases current health up to a maximum of max health
func (p *Player) kevJuice(Health int, Mana int, Stamina int) {
	p.health += 100
	p.mana += 100
	p.stamina += 50
	fmt.Printf("You drank some kevJuice! You feel revitalized")
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	if p.stamina > p.maxStamina {
		p.stamina = p.maxStamina
	}
	if p.mana > p.maxMana {
		p.mana = p.maxMana
	}
}

// Method for decreasing the player's health if an enemy's swing connects
func (p *Player) gotHurt(Health int) {
	hurt := true
	if hurt {
		p.health -= 25
		fmt.Printf("Ouch! You have %d health remaining\n", p.health)
	} else {
		fmt.Printf("A near miss, you're unscathed!\n")
	}
}

// Method for casting fireball, if successful, will cost mana
func (p *Player) fireball(Mana int) {
	fizzle := false
	if fizzle {
		fmt.Printf("Oh no! Your spell fizzled!\n")
	} else {
		p.mana -= 50
		fmt.Printf("BOOM! Your enemy is burning, you have %d mana remaining\n", p.mana)
	}
}

// Method for a Regular swing, costing stamina hit or miss
func (p *Player) swing(Stamina int) {
	hit := true
	p.stamina -= 10
	if hit {
		fmt.Printf("A mighty strike! You have %d stamina left\n", p.stamina)
	} else {
		fmt.Printf("You whiffed! You have %d stamina left\n", p.stamina)
	}
}

// Method that levels the player up, increasing their level count and stats
func (p *Player) levelUp() {
	p.xp = 0
	p.xpTillLevel += 50
	fmt.Printf("You leveled up! You gained 50 maxHP, 50 maxMP, and 10 maxSP\n")
	p.maxHealth += 50
	p.maxMana += 50
	p.maxStamina += 10
	p.level += 1
}

// Method to list off player1's stats
func (p *Player) Status() {
	//The real TODO here is to make it format nicer on the console. Use ----- and | etc. (Put it in a box)
	//Look into the fmt documentation on how to have %d be fixed width/fixed number of characters/left/right justified.
	//make the first argument of this printf a constant.
	fmt.Printf(`Character: %s
Level: %d
Experience %d/%d
HP: %d out of %d
MP: %d out of %d
SP: %d out of %d

`,
		p.name,
		p.level,
		p.xp, p.xpTillLevel,
		p.health, p.maxHealth,
		p.mana, p.maxMana,
		p.stamina, p.maxStamina,
	)

}

// Method that rewards the player when a monster is defeated
func (p *Player) killedMonster(xp int, gold int) {
	victory := true
	if victory {
		p.xp += 100
		p.gold += 100
		if p.xp >= p.xpTillLevel {
			p.levelUp()
		}
	}
}
