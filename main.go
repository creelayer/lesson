package main

import (
	"fmt"
	"github.com/na1al/lesson/pkg/geo"
	"github.com/na1al/lesson/pkg/address"
)

func main()  {
	l := geo.Location{}
	l.Id = 1;

	a := address.Adr{2}

	fmt.Print("commit 1")

	fmt.Print(a.Name)
}
