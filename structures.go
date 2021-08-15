package main

import "fmt"
//Go does not have classes in it. Instead we use structures and methods can be placed within structures for usage.

//Example structure for a car and its properties would be
type car struct {
	gasPedal      uint16 //min: 0,      max: 65535
	brakeSpeed    uint16 //Same range as the variable above
	steeringWheel int16  //min: -32768  max: 32768
	topSpeedKmph  float64
}

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmph/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmph/kmh_multiple/usixteenbitmax)
}
// This is a pointer receiver that receives a pointer struct and modifies the value of the struct object (car in this
// case). In brief, func below modifies the struct itself via pointer type.
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmph = newSpeed
}

func main() {
	//aCar is a car structure and can be declared as shown below

	//aCar := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	//aCar := car{gas_pedal_pos,brake_pedal_pos,steering_wheel_pos,top_speed_kmh}
	aCar := car{22314, 0, 12562, 225.0}

	fmt.Println("gas_pedal:", aCar.gasPedal, "brake_pedal:", aCar.brakeSpeed, "steering_wheel:", aCar.steeringWheel,
		"top_speed:", aCar.topSpeedKmph)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(500)
	fmt.Println("After changes made with the pointer receiver function: ")
	fmt.Println("gas_pedal:", aCar.gasPedal, "brake_pedal:", aCar.brakeSpeed, "steering_wheel:", aCar.steeringWheel,
		"top_speed:", aCar.topSpeedKmph)
}