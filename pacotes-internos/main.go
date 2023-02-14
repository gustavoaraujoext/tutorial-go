package main

import (
	"fmt"

	"github.com/kustavo/tutorial-go/pacotes-internos/d"
	"github.com/kustavo/tutorial-go/pacotes-internos/internal/a"
	"github.com/kustavo/tutorial-go/pacotes-internos/internal/c"
	"github.com/kustavo/tutorial-go/pacotes-internos/d/e"
)

func main() {
	fmt.Println("Método Main")
	a.A("Main")
	c.C("Main")
	d.D("Main")
	e.E("Main")
}
