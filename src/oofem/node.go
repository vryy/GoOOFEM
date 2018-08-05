package oofem

import "fmt"

type Node struct {
	mPosition     [3]float64
	mDisplacement [3]float64
	mConstraint   *Constraint
	mForce        *Force
	mDofNumbers   [3]int
}

func NodeXYZ(x float64, y float64, z float64) *Node {
	node := Node{
		mPosition:     [3]float64{x, y, z},
		mDisplacement: [3]float64{0, 0, 0},
		mConstraint:   ConstraintXYZ(true, true, true),
		mForce:        ForceXYZ(0, 0, 0),
		mDofNumbers:   [3]int{-1, -1, -1},
	}
	return &node
}

func (node *Node) SetConstraint(c *Constraint) {
	node.mConstraint = c
}

func (node Node) GetConstraint() *Constraint {
	return node.mConstraint
}

func (node *Node) SetForce(f *Force) {
	node.mForce = f
}

func (node Node) GetForce() *Force {
	return node.mForce
}

func (node *Node) EnumerateDofs(start int) int {
	for i := 0; i < 3; i++ {
		if node.mConstraint.IsFree(i) {
			node.mDofNumbers[i] = start
			start = start + 1
		} else {
			node.mDofNumbers[i] = -1
		}
	}
	return start
}

func (node Node) GetDofNumbers() [3]int {
	return node.mDofNumbers
}

func (node Node) GetPosition() [3]float64 {
	return node.mPosition
}

func (node *Node) SetDisplacement(u [3]float64) {
	node.mDisplacement[0] = u[0]
	node.mDisplacement[1] = u[1]
	node.mDisplacement[2] = u[2]
}

func (node Node) GetDisplacement() [3]float64 {
	return node.mDisplacement
}

func (node Node) Print() {
	fmt.Printf("%f\t%f\t%f\n", node.mPosition[0], node.mPosition[1], node.mPosition[2])
}
