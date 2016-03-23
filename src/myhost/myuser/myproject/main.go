package main

import (
	"fmt"
	"origA/userA/projA"
	"origB/userB/projB"
)

func main() {
	A := projA.New()
	fmt.Println("A comes from", A.Location)
	B := projB.New()
	fmt.Println("B.A comes from", B.A.Location)
}
