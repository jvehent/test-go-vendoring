package main

import (
	"fmt"
	"origB/userB/projB"
)
import "origA/userA/projA"

func main() {
	A := projA.New()
	fmt.Println("A comes from", A.Location)
	B := projB.New()
	fmt.Println("B.A comes from", B.A.Location)
}
