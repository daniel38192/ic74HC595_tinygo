package ic74hc595

import (
	"machine"
	"time"
)

// States struct represents the states of individual output pins of the 74HC595 shift register.
type States struct {
	I1 bool //QH
	I2 bool //QG
	I3 bool //QF
	I4 bool //QE
	I5 bool //QD
	I6 bool //QC
	I7 bool //QB
	I8 bool //QA
}

// IC74HC595A struct represents the 74HC595 shift register and its control pins.
type IC74HC595A struct {
	dirSer   machine.Pin // Serial Data pin
	dirClk   machine.Pin //Clock Pin
	dirLatch machine.Pin //Data Pin
}

var defaultOff = States{
	I1: false,
	I2: false,
	I3: false,
	I4: false,
	I5: false,
	I6: false,
	I7: false,
	I8: false,
}

// New initializes a new IC74HC595A instance with specified pin configurations.
func New(dirSer machine.Pin, dirClk machine.Pin, dirLatch machine.Pin) IC74HC595A {
	ic := IC74HC595A{
		dirSer:   dirSer,
		dirClk:   dirClk,
		dirLatch: dirLatch,
	}
	ic.Configure()
	return ic
}

// Configure configures the control pins as output pins.
func (ic IC74HC595A) Configure() {
	ic.dirSer.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ic.dirClk.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ic.dirLatch.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// sendData sends data to the 74HC595 shift register.
func (ic IC74HC595A) SendData(args States) {
	// Sending data to serial pin sequentially
	ic.dirSer.Set(args.I1)
	ic.clockWrite()
	ic.dirSer.Set(args.I2)
	ic.clockWrite()
	ic.dirSer.Set(args.I3)
	ic.clockWrite()
	ic.dirSer.Set(args.I4)
	ic.clockWrite()
	ic.dirSer.Set(args.I5)
	ic.clockWrite()
	ic.dirSer.Set(args.I6)
	ic.clockWrite()
	ic.dirSer.Set(args.I7)
	ic.clockWrite()
	ic.dirSer.Set(args.I8)
	ic.clockWrite()
}

// Show updates the outputs of the 74HC595 shift register.
func (ic IC74HC595A) Show() {
	time.Sleep(time.Microsecond)
	ic.dirLatch.High()
	time.Sleep(time.Microsecond)
	ic.dirLatch.Low()
	time.Sleep(time.Microsecond)
}

// Clear clears the outputs of the 74HC595 shift register.
func (ic IC74HC595A) Clear() {
	ic.SendData(defaultOff)
	ic.Show()
}

// clockWrite toggles the clock pin to shift data into the shift register.
func (ic IC74HC595A) clockWrite() {
	time.Sleep(time.Microsecond)
	ic.dirClk.High()
	time.Sleep(time.Microsecond)
	ic.dirClk.Low()
	time.Sleep(time.Microsecond)

}
