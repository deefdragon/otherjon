package main

type Car struct {
	wheeles []Wheele
	doors   []Door
	vin     string
	license string

	seats []Seat
}

type Seat struct {
	isOccupied bool
	//
}

type Wheele struct {
	axelNumber   int
	side         string
	miles        float64
	wheeleDamage Damage
}

type Damage struct {
	//stuff related to the damage a thing has suffered.
}

type Door struct {
	orderNumber int
	side        string
	doorDamage  Damage
}

func calculateCarWorth(thingamajig Car, number int) float64 {
	thingamajig.CalculateWorth()
	return 0

}

//method
func (car *Car) CalculateWorth() float64 {
	return 0
}

type credicCheck struct {
	//
	currentChecCount int
}

func calculateCredit(c credicCheck) int {
	c.currentChecCount++
	return 0
}

func (c *credicCheck) calculateCredit() int {
	c.currentChecCount++
	x := 5
	x2(&x)
	//x == 50
	return 0

}

func x2(in *int) int {
	*in = *in * 10
	return *in
}
