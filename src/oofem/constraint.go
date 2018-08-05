package oofem

import "fmt"

type Constraint struct {
	mFree [3]bool
}

func ConstraintXYZ(c1 bool, c2 bool, c3 bool) *Constraint {
	nc := new(Constraint)
	nc.mFree[0] = c1
	nc.mFree[1] = c2
	nc.mFree[2] = c3
	return nc
}

func (c Constraint) IsFree(i int) bool {
	return c.mFree[i]
}

func (c Constraint) Print() {
	fmt.Printf("%t\t%t\t%t\n", c.IsFree(0), c.IsFree(1), c.IsFree(2))
}
