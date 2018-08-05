package main

import (
	"../src/oofem"
	"fmt"
)

func main() {
	c1 := oofem.ConstraintXYZ(true, false, true)

	f1 := oofem.ForceXYZ(1.2, -4, 0)

	n1 := oofem.NodeXYZ(1, 0, 0)
	n2 := oofem.NodeXYZ(0, 1, 0)

	n1.SetConstraint(c1)
	n2.SetForce(f1)

	e1 := oofem.LinearTrussElement(2.1e8, 0.2, n1, n2)

	fmt.Printf("Constraint (u1, u2, u3)\n")
	n1.GetConstraint().Print()

	fmt.Printf("Force (r1, r2, r3)\n")
	n2.GetForce().Print()

	fmt.Printf("Nodes (x1, x2, x3)\n")
	n1.Print()
	n2.Print()

	fmt.Printf("Element (E, A, L)\n")
	e1.Print()
}
