package projA

type A struct {
	Location string
}

func New() A {
	return A{Location: "$GOPATH/src/myhost/myuser/myproject/vendor/origB/userB/projB/vendor/origA/userA/projA/"}
}
