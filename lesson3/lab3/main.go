package main

import "player"

func main() {
	var player1 = player.PlayerData{
		name:        "Jon",
		level:       1,
		xp:          0,
		maxHealth:   100,
		maxMana:     100,
		maxStamina:  50,
		xpTillLevel: 100,
		gold:        0,
		Health:      100,
		Mana:        100,
		Stamina:     50,
	}

	player1.pStatus()
	player1.gotHurt(player1.Health)
	player1.fireball(player1.Mana)
	player1.fireball(player1.Mana)
	player1.swing(player1.Stamina)
	player1.killedMonster(player1.xp, player1.gold)
	player1.kevJuice(player1.Health, player1.Mana, player1.Stamina)
	player1.pStatus()
}
