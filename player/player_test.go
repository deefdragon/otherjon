package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
	p := New("foo")
	if p.Name != "foo" {
		t.Fail()
	}
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
