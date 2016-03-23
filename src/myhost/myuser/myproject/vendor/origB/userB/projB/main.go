package projB

import "origA/userA/projA"

type B struct {
	A projA.A
}

func New() B {
	return B{A: projA.New()}
}
