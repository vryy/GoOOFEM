package oofem

type Node struct {
	mPosition     [3]double
	mDisplacement [3]double
	mConstraint   Constraint
	mForce        Force
}
