package player

import "fmt"

type Player struct {
	//TODO Getname,
	Name string

	//TODO Getlevel,
	Level int

	//TODO GetXp, AddXp
	Xp int

	//TODO Gethealth, ChangeHealth, IsDead
	Health int

	//TODO Getmana, ChangeMana, CanCast
	Mana int

	//TODO Getstamina, ChangeStamina
	Stamina int

	//TODO GetmaxHealth,
	MaxHealth int

	//TODO GetmaxMana,
	MaxMana int

	//TODO GetmaxStamina,
	MaxStamina int

	//TODO GetxpTillLevel,
	XpTillLevel int

	//TODO Getgold, ChangeGold, HasGold
	Gold int

	//TODO print player stats to console (Status)

	//TODO remove all unnecessary functions.
}

func (p *Player) GetName() string {
	return p.Name
}
func (p *Player) GetLevel() int {
	return p.Level
}
func (p *Player) GetXp() int {
	return p.Xp
}
func (p *Player) AddXp(x int) {
	p.Xp += x
	if p.Xp <= 0 {
		p.Xp = 0
	}
	if p.Xp >= p.XpTillLevel {
		p.Xp -= p.XpTillLevel
		p.levelUp()
		p.AddXp(0)
	}
}
func (p *Player) GetHealth() int {
	return p.Health
}
func (p *Player) ChangeHealth(h int) {
	p.Health += h
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}
func (p *Player) IsDead() {
	if p.Health <= 0 {
		p.Health = 0
		p.Name = ("{Deceased}" + p.Name)
	}
}
func (p *Player) GetMana() int {
	return p.Mana
}
func (p *Player) ChangeMana(m int) {
	p.Mana += m
	if p.Mana > p.MaxMana {
		p.Mana = p.MaxMana
	}
	if p.Mana < 0 {
		p.Mana = 0
	}
}
func (p *Player) CanCast(cost int) bool {
	if cost > p.Mana {
		return false
	}
	return true
}
func (p *Player) GetStamina() int {
	return p.Stamina
}
func (p *Player) ChangeStamina(s int) {
	p.Stamina += s
	if p.Stamina > p.MaxStamina {
		p.Stamina = p.MaxStamina
	}
	if p.Stamina < 0 {
		p.Stamina = 0
	}
}
func (p *Player) CanSwing(cost int) bool {
	if cost > p.Stamina {
		return false
	}
	return true
}
func (p *Player) GetMaxHealth() int {
	return p.MaxHealth
}
func (p *Player) GetMaxMana() int {
	return p.MaxMana
}
func (p *Player) GetMaxStamina() int {
	return p.MaxStamina
}
func (p *Player) GetXpTillLevel() int {
	return p.XpTillLevel
}
func (p *Player) GetGold() int {
	return p.Gold
}
func (p *Player) ChangeGold(g int) {
	p.Gold += g
	if p.Gold < 0 {
		p.Gold = 0
	}
}
func (p *Player) HasGold(cost int) bool {
	if cost <= p.Gold {
		return true
	}
	return false
}

func New(name string) *Player {

	return &Player{
		Name: name,
		Xp:   0, XpTillLevel: 100,
		Level:  1,
		Health: 100, MaxHealth: 100,
		Mana: 100, MaxMana: 100,
		Stamina: 50, MaxStamina: 50,
		Gold: 0,
	}
}

// Method that levels the player up, increasing their level count and stats
func (p *Player) levelUp() {
	p.XpTillLevel += 50
	fmt.Printf("You leveled up! You gained 50 maxHP, 50 maxMP, and 10 maxSP\n")
	p.MaxHealth += 50
	p.MaxMana += 50
	p.MaxStamina += 10
	p.Level += 1
}

// Method to list off player1's stats
func (p *Player) Status() {
	//The real TODO here is to make it format nicer on the console. Use ----- and | etc. (Put it in a box)
	//Look into the fmt documentation on how to have %d be fixed width/fixed number of characters/left/right justified.
	//make the first argument of this printf a constant.
	fmt.Printf(`Character: %s
|==============|
|Level: %2d     |
|XP: %3d / %-3d |
|HP: %3d / %-3d |
|MP: %3d / %-3d |
|SP: %3d / %-3d |
|==============|
`,
		p.Name,
		p.Level,
		p.Xp, p.XpTillLevel,
		p.Health, p.MaxHealth,
		p.Mana, p.MaxMana,
		p.Stamina, p.MaxStamina,
	)
}
