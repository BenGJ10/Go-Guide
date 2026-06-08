package main

import "fmt"

type socket struct{}

type PowerDrawer interface {
	Draw(power int)
}

func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}

type Mobile struct {
	brand string
}

func (m Mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type Laptop struct {
	cpu string
}

func (l Laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

func main() {
	socket := socket{}

	mobile := Mobile{"Apple"}
	socket.Plug(mobile, 100)

	laptop := Laptop{"Intel"}
	socket.Plug(laptop, 200)
}
