package projA

type A struct {
	Location string
}

func New() A {
	return A{Location: "$GOPATH/src/origA/userA/projA"}
}
