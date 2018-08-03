package oofem

import "fmt"

type Constraint struct {
	mFree [3]bool
}

func (c Constraint) isFree(i int) bool {
	return c.mFree[i]
}

func (c Constraint) print() {
	fmt.Print("Constraint: %d %d %d\n", c.isFree(0), c.isFree(1), c.isFree(2))
}
