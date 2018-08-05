package oofem

import (
	"fmt"
	"math"
)

type TrussElement struct {
	mNode1      *Node
	mNode2      *Node
	mEModulus   float64
	mArea       float64
	mDofNumbers [6]int
}

func LinearTrussElement(E float64, A float64, n1 *Node, n2 *Node) *TrussElement {
	ne := TrussElement{
		mNode1:      n1,
		mNode2:      n2,
		mEModulus:   E,
		mArea:       A,
		mDofNumbers: [6]int{-1, -1, -1, -1, -1, -1},
	}
	return &ne
}

func (elem TrussElement) GetNode1() *Node {
	return elem.mNode1
}

func (elem TrussElement) GetNode2() *Node {
	return elem.mNode2
}

func (elem *TrussElement) EnumerateDofs() {
	elem.mDofNumbers[0] = elem.GetNode1().GetDofNumbers()[0]
	elem.mDofNumbers[1] = elem.GetNode1().GetDofNumbers()[1]
	elem.mDofNumbers[2] = elem.GetNode1().GetDofNumbers()[2]
	elem.mDofNumbers[3] = elem.GetNode2().GetDofNumbers()[0]
	elem.mDofNumbers[4] = elem.GetNode2().GetDofNumbers()[1]
	elem.mDofNumbers[5] = elem.GetNode2().GetDofNumbers()[2]
}

func (elem TrussElement) GetDofNumbers() [6]int {
	return elem.mDofNumbers
}

func (elem TrussElement) GetArea() float64 {
	return elem.mArea
}

func (elem TrussElement) GetEModulus() float64 {
	return elem.mEModulus
}

func (elem TrussElement) GetLength() float64 {
	L2 := 0.0
	for i := 0; i < 3; i++ {
		L2 += math.Pow(elem.GetNode2().GetPosition()[i]-elem.GetNode1().GetPosition()[i], 2)
	}
	return math.Sqrt(L2)
}

func (elem TrussElement) Print() {
	fmt.Printf("%.6e\t%.6e\t%.6e", elem.GetEModulus(), elem.GetArea(), elem.GetLength())
}
