package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLevel(t *testing.T) {
	p := New("name")
	assert.Equal(t, 1, p.GetLevel())
}
func TestGetXp(t *testing.T) {
	p := New("name2")
	assert.Equal(t, 0, p.GetXp())
}
func TestGetHealth(t *testing.T) {
	p := New("name3")
	assert.Equal(t, 100, p.GetHealth())
}
func TestChangeHealth(t *testing.T) {
	p1 := New("name4")
	p1.AddXp(250)
	p1.ChangeHealth(101)
	p1.GetStatus()
	assert.Equal(t, 200, p1.Health)
	assert.Equal(t, 200, p1.MaxHealth)
}
func TestGetMana(t *testing.T) {
	p := New("name5")
	assert.Equal(t, 100, p.GetMana())
}
func TestChangeMana(t *testing.T) {
	p := New("name6")
	p.ChangeMana(-25)
	assert.Equal(t, 75, p.Mana)
	p.ChangeMana(50)
	assert.Equal(t, 100, p.Mana)
	p.ChangeMana(-9001)
	assert.Equal(t, 0, p.Mana)
}
func TestCanCast(t *testing.T) {
	p := New("name7")
	assert.False(t, false, p.CanCast(9001), "Player could cast without having mana")
	assert.True(t, true, p.CanCast(100), "Player could not cast, even with enough mana")
}
func TestGetStamina(t *testing.T) {
	p := New("name8")
	assert.Equal(t, 50, p.GetStamina())
}
func TestChangeStamina(t *testing.T) {
	p := New("name9")
	p.ChangeStamina(-10)
	assert.Equal(t, 40, p.Stamina)
	p.AddXp(250)
	p.ChangeStamina(50)
	assert.Equal(t, 70, p.Stamina)
	p.ChangeStamina(-9001)
	assert.Equal(t, 0, p.Stamina)
}
func TestCanSwing(t *testing.T) {
	p := New("name10")
	assert.True(t, p.CanSwing(10))
	assert.False(t, p.CanSwing(100))
}
func TestCreatePlayer(t *testing.T) {
	p := New("foo")
	if p.Name != "foo" {
		t.Fail()
	}
}
func TestGetMaxHealth(t *testing.T) {
	p := New("name™")
	assert.Equal(t, 100, p.GetMaxHealth())
}
func TestGetMaxMana(t *testing.T) {
	p := New("name™")
	assert.Equal(t, 100, p.GetMaxMana())
}
func TestGetMaxStamina(t *testing.T) {
	p := New("name™")
	assert.Equal(t, 50, p.GetMaxStamina())
}
func TestGetXpTillLevel(t *testing.T) {
	p := New("name™")
	assert.Equal(t, 100, p.GetXpTillLevel())
}
func TestGetGold(t *testing.T) {
	p := New("name™")
	assert.Equal(t, 10, p.GetGold())
}
func TestChangeGold(t *testing.T) {
	p := New("Me")
	p.ChangeGold(-100)
	assert.Equal(t, 0, p.Gold)
}
func TestExp(t *testing.T) {
	p := New("")
	p.AddXp(10)
	if p.Xp != 10 {
		t.Fail()
	}
	if p.XpTillLevel != 100 {
		t.Fail()
	}
	if p.Level != 1 {
		t.Fail()
	}
}

func TestLeveling(t *testing.T) {
	p := New("")
	p.AddXp(100)
	if p.Xp != 0 {
		t.Fail()
	}
	if p.XpTillLevel != 150 {
		t.Fail()
	}
	if p.Level != 2 {
		t.Fail()
	}
	if p.MaxHealth != 150 {
		t.Fail()
	}
	if p.MaxMana != 150 {
		t.Fail()
	}
	if p.MaxStamina != 60 {
		t.Fail()
	}
}

// func statusHelper(name string, current int, max int) string {
// 	fmt.Printf("|%s: %3d / %-3d |\n", name, current, max)
// 	return ""
// }
func TestGetStatus_Starting(t *testing.T) {
	p := New("SansUndertale")
	const GetStatus_Starting1 = `Character: SansUndertale
|==============|
|Level:  1     |
|XP:   0 / 100 |
|HP: 100 / 100 |
|MP: 100 / 100 |
|SP:  50 / 50  |
|==============|
`
	const GetStatus_Starting2 = `Character: SansUndertale
|==============|
|Level:  3     |
|XP:   1 / 200 |
|HP: 100 / 200 |
|MP: 100 / 200 |
|SP:  50 / 70  |
|==============|
`
	assert.Equal(t, GetStatus_Starting1, p.GetStatus())
	p.AddXp(251)
	assert.Equal(t, GetStatus_Starting2, p.GetStatus())
}

func TestMultiLevelLeveling(t *testing.T) {
	p := New("")
	p.AddXp(1050)

	assert.Equal(t, 50, p.Xp, "XP Calculation failed")
	assert.Equal(t, 350, p.XpTillLevel, "XP Until Level Calculation failed")
	assert.Equal(t, 6, p.Level, "Level Calculation failed")

}

func TestIsDeadNameUpdate(t *testing.T) {
	p := New("foo")
	p.ChangeHealth(-500)
	assert.True(t, p.IsDead(), "Player was not dead")
	assert.Equal(t, p.GetName(), "foo {dead}", "name did not update to notify deceased status")
	assert.True(t, p.IsDead(), "Player was not dead")
	assert.Equal(t, p.GetName(), "foo {dead}", "name did not update to notify deceased status")
}

func TestIsDeadWhenNotDeadNameUpdate(t *testing.T) {
	p := New("foo")
	assert.False(t, p.IsDead(), "Player is not dead")
	assert.Equal(t, p.GetName(), "foo", "name was changed")
}

func TestHasGold(t *testing.T) {
	p := New("foo")
	assert.True(t, p.HasGold(9))
	assert.True(t, p.HasGold(10))
	assert.False(t, p.HasGold(11))

	assert.True(t, p.HasGold(0))
	assert.True(t, p.HasGold(-10))
	assert.False(t, p.HasGold(100))

}
