package oofem

import "fmt"

type Force struct {
	mComponent [3]double
}

func New() Force {
	f = new(Force)
	return f
}

func (f Force) getComponent(i int) double {
	return f.mComponent[i]
}

func (f Force) setComponent(i int, v double) {
	f.mComponent[i] = v
}

func (f Force) print() {
	fmt.Print("Force: %f %f %f\n", f.getComponent(0), f.getComponent(1), f.getComponent(2))
}
