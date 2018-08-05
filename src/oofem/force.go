package oofem

import "fmt"

type Force struct {
	mComponent [3]float64
}

func ForceXYZ(f1 float64, f2 float64, f3 float64) *Force {
	nf := new(Force)
	nf.mComponent[0] = f1
	nf.mComponent[1] = f2
	nf.mComponent[2] = f3
	return nf
}

func (f Force) GetComponent(i int) float64 {
	return f.mComponent[i]
}

func (f Force) SetComponent(i int, v float64) {
	f.mComponent[i] = v
}

func (f Force) Print() {
	fmt.Printf("%.6f\t%.6f\t%.6f\n", f.GetComponent(0), f.GetComponent(1), f.GetComponent(2))
}
