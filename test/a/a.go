package a

import (
	"fmt"
	"pervasive-chain/test/i"
)

type PackageA struct {
	B i.PackageBInterface
}

func (a PackageA) PrintA() {
	fmt.Println("I'm a!")
}

func (a PackageA) PrintAll() {
	a.PrintA()
	a.B.PrintB()
}