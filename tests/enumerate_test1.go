package main

import (
	"fmt"
	"oofem"
)

func main() {
	eqn := 0

	n1 := oofem.NodeXYZ(0, 0, 0)
	n2 := oofem.NodeXYZ(1, 1, 0)

	e1 := oofem.LinearTrussElement(1, 1, n1, n2)

	c1 := oofem.ConstraintXYZ(false, true, true)

	n1.SetConstraint(c1)

	eqn = n1.EnumerateDofs(eqn)
	eqn = n2.EnumerateDofs(eqn)

	e1.EnumerateDofs()

	fmt.Printf("Number of equations: %d\n", eqn)
	fmt.Printf("Equation numbers of nodal DOFs\n")
	fmt.Printf("%d\n", n1.GetDofNumbers())
	fmt.Printf("%d\n", n2.GetDofNumbers())
	fmt.Printf("Equation numbers of element DOFs\n")
	fmt.Printf("%d\n", e1.GetDofNumbers())
}
