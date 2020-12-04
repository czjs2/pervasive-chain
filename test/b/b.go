package b

import (
	"fmt"
	"pervasive-chain/test/i"
)

type PackageB struct {
	A i.PackageAInterface
}

func (b PackageB) PrintB() {
	fmt.Println("I'm b!")
}

func (b PackageB) PrintAll() {
	b.PrintB()
	b.A.PrintA()
}
